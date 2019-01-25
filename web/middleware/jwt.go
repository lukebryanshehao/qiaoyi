package middleware

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"io"
	"io/ioutil"
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
	"strings"
	"time"
)

//注册jwt中间件
func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			fmt.Println("----------token:",token.Raw)

			//自己加密的秘钥或者说盐值
			return []byte("My spacej"), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)
		ErrorHandler: func(ctx iris.Context, s string) {



			if strings.Contains(ctx.Request().RequestURI, "/material/generate/word") {
				ctx.Next()
			}
			if strings.Contains(ctx.Request().RequestURI, "/meeting/review/recommend/template/download") {
				ctx.Next()
			} else {
				fmt.Println("错误打印",s)
				result := model.ResultBean{Status:false, Msg:"认证失败，请重新登录认证"}
				ctx.JSON(result)

			}
		},
	})
	return jwtHandler
}

//生成token
func GenerateToken(user *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username, //用户信息
		"session":  user.Session, //session
		"id":         user.ID,   //用户信息
		"usertypeid": user.UserTypeId,
		"iss":        "Iris",                                                   //签发者
		"iat":        time.Now().Unix(),                                        //签发时间
		"jti":        "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp":        time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})
	tokenString, _ := token.SignedString([]byte("My spacej"))
	fmt.Println("签发时间：",time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10 * time.Hour * time.Duration(1)).Unix())
	return tokenString
}

func ParseToken(tokenString string, key string) (interface{}, bool){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func CheckSession(ctx iris.Context) {
	fmt.Println("请求接口:",ctx.Request().URL)

	token := ctx.GetHeader("Authorization")
	if token != "" && len(token)>7 {
		token = token[7 : len(token)]
	}
	fmt.Println(token)
	var sessionID = ""
	if token != "" && token != "undefined" {
		v,_ := ParseToken(token, "My spacej")
		sessionID = v.(jwt.MapClaims)["session"].(string)
	}

	fmt.Println("---------sessionID:----------------",sessionID)
	var onlineSessionIDList = SMgr.GetSessionIDList()
	for _, onlineSessionID := range onlineSessionIDList {
		fmt.Println("在线用户SessionID:",onlineSessionID)
		if onlineSessionID == sessionID {
			ctx.Next()
			return
		}
	}
	fmt.Println("接口",ctx.Request().URL,"-------------------被拦截了")

	fmt.Println("-------66666666666--------",ctx.GetHeader("Authorization"))

	result := model.ResultBean{Status:false,Code:"512", Msg:"该账号已在其它设备登陆!"}
	ctx.JSON(result)
	//ctx.Next()
}

func SaveSystemLog(ctx iris.Context)  {
	url := ctx.Request().URL
	fmt.Println("请求接口:",url.String())
	var body io.ReadCloser
	body = ctx.Request().Body
	//ReadAll时会close body,需要重新设置body内容
	bodyStr, _ := ioutil.ReadAll(body)
	fmt.Println("请求参数:",string(bodyStr))
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyStr))
	var nickName string
	var userID uint
	token := ctx.GetHeader("Authorization")
	if token != "" && len(token)>7 {
		token = token[7 : len(token)]
	}
	if token != "" && token != "undefined" {
		v,_ := ParseToken(token, "My spacej")
		sessionID := v.(jwt.MapClaims)["session"].(string)
		if userInfo, ok := SMgr.GetSessionVal(sessionID, "UserInfo"); ok {
			if value, ok := userInfo.(model.User); ok {
				nickName = value.Username
				userID = value.ID
			}
		}
	}
	fmt.Println("token-----",token)
	fmt.Println("Username-----",nickName)
	fmt.Println("userID-----",userID)

	if (token != "" && token != "undefined") || url.String() == "/user/login" {
		var systemLog model.SystemLog
		if url.String() == "/user/login" {
			var user model.User
			//{"loginname":"admin2018","password":"123456"}
			bodyStrArr := strings.Split(string(bodyStr), "\"")
			nickName = bodyStrArr[3]
			if userID == 0 && nickName != ""{
				datasource.DB.Where("login_name = ?", nickName).First(&user)
				userID = user.ID
			}
		}
		systemLog.LoginName = nickName
		systemLog.UserId = userID

		operate := ""

		if url.String() == "/user/login" {
			operate = "登陆"
		}else {
			for key,value := range ControllerMap{
				if strings.HasPrefix(url.String(), key) {
					operate = value
					break
				}
			}
		}

		systemLog.Operate = operate
		systemLog.Method = url.String()
		systemLog.Parameter = string(bodyStr)
		systemLog.Ip = ctx.RemoteAddr()
		if operate != "" && userID != 0 {
			go SaveLog(systemLog)
		}
	}
	ctx.Next()
}

func SaveLog(systemLog model.SystemLog)  {
	tran := datasource.DB.Begin()
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err == nil {
			//提交事务
			tran.Commit()
		} else {
			//回滚
			tran.Rollback()
		}
	}()

	if err := tran.Save(&systemLog).Error; err != nil{
		fmt.Println(err)
	}
}