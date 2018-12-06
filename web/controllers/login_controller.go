package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
	"qiaoyi_back/utils"
	"time"
)

type LoginController struct {
	Service services.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{Service: services.NewLoginService(repositorys.NewLoginRepository())}
}

var key = "nKhiUoaSwRegTviVohjAOd"

func (c *LoginController) Get(context iris.Context) (mvc.Result)  {
	return mvc.View{
		Name: "login.html",
		Data: nil,
	}
}

func (c *LoginController) PostLogin(context iris.Context) (mvc.Result) {
	user := &model.User{}
	user.Username = context.Request().FormValue("username")
	user.Password = context.Request().FormValue("password")
	user.Checkcode = context.Request().FormValue("checkcode")
	//context.ReadJSON(&user)

	var tokenState string
	var tokenString string

	if user.Username == "" || user.Password == "" {
		cookie, err := context.Request().Cookie("token")
		if err != nil {
			fmt.Println("获取cookie错误或者没有cookie")
			tokenState = "获取cookie错误或者没有cookie"
		}else {
			//cookie.Value
			if cookie.Value != "" {
				tokenString = cookie.Value
			}

			claims, ok := utils.ParseToken(tokenString, key)
			if ok {
				// 时间在u 之前
				//is_after := t.After(t_new)
				oldT, _ := time.Parse("2006-01-02 15:04:05",claims.(jwt.MapClaims)["exp"].(string))
				ct := time.Now()
				if  ct.Before(oldT){
					ok = false
					tokenState = "Token 已过期,请重新登陆"

				} else {
					tokenState = "Token 正常"
				}

				var username = claims.(jwt.MapClaims)["username"].(string)
				var password = claims.(jwt.MapClaims)["password"].(string)
				user.Username = username
				user.Password = password

			}else {
				tokenState = "Token 无效,登陆失败!"
			}
		}

	}

	exist,userInfo:= c.Service.Exist(user)

	if exist {
		if tokenState != "Token 正常" {
			t := time.Now()
			type UserInfo map[string] interface{}
			userInfo := make(UserInfo)
			userInfo["username"] = user.Username
			userInfo["password"] = user.Password
			userInfo["exp"] = t.AddDate(0,2,0)

			tokenString = utils.CreateToken(key,userInfo)

			COOKIE_MAX_MAX_AGE := time.Hour * 24 * 30 *2/ time.Second   // 单位：秒。
			maxAge := int(COOKIE_MAX_MAX_AGE)

			token_cookie:=&http.Cookie{
				Name:   "token",
				Value:    tokenString,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   maxAge,
			}
			context.SetCookie(token_cookie)
		}
		tokenState = "登陆成功!"
	}else {
		tokenState = "用户名或密码错误,登陆失败!"
	}

	html := "login.html"
	resultBean := model.CreateResultWithMsg(tokenState)
	if tokenState == "登陆成功!" {
		html = "index-2.html"
		resultBean = model.CreateResultWithData(userInfo)
	}
	view := mvc.View{
		Name: html,
		Data: resultBean,
	}
	return view
}

func (c *LoginController) PostLoginout(context iris.Context) {
	context.RemoveCookie("token")
	//如果要设置自定义路径：
	// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))
}

func (c *LoginController) PostUpdatepassword(context iris.Context) (model.ResultBean)  {
	return model.CreateResultWithMsg("")
}

func (c *LoginController) PostGetinfo(context iris.Context) {
	user := c.Service.GetInfo(context.Request().FormValue("username"))
	resultBean := model.CreateResultWithMsg("h获取失败")
	if user.ID != 0{
		resultBean = model.CreateResultWithData(user)
	}
	json, err := json.Marshal(resultBean)
	if err != nil {
		panic(err)
	}
	context.WriteString(string(json))
}



