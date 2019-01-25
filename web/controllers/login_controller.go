package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
	"qiaoyi_back/web/middleware"
	"time"
)

type LoginController struct {
	Service services.LoginService
	Ctx iris.Context
}

func NewLoginController() *LoginController {
	return &LoginController{Service: services.NewLoginService(repositorys.NewLoginRepository())}
}

func (c *LoginController) Get() (mvc.Result)  {
	return mvc.View{
		Name: "login.html",
		Data: nil,
	}
}

func (c *LoginController) PostLogin() *model.ResultBean {
	var user model.User
	c.Ctx.ReadJSON(&user)
	//user.Username = c.Ctx.Request().FormValue("username")
	//user.Password = c.Ctx.Request().FormValue("password")
	if user.Username == "" || user.Password == "" {
		return model.NewResultBean(false, "用户名或密码不能为空")
	}

	var token string

	exist:= c.Service.Exist(&user)
	if exist {
		//创建客户端对应cookie以及在服务器中进行记录
		var sessionID = middleware.SMgr.StartSession(c.Ctx.ResponseWriter(), c.Ctx.Request())
		fmt.Println("-------------------创建新的sessionID:",sessionID)
		//192.168.0.115
		ip := c.Ctx.RemoteAddr()
		user.RemoteAddr = ip
		user.Session = sessionID
		user.AccessTime = time.Now()
		var loginUserInfo =	user

		//踢除重复登录的
		var onlineSessionIDList = middleware.SMgr.GetSessionIDList()
		for _, onlineSessionID := range onlineSessionIDList {
			fmt.Println("-------------------onlineSessionID:",onlineSessionID)
			if userInfo, ok := middleware.SMgr.GetSessionVal(onlineSessionID, "UserInfo"); ok {
				if value, ok := userInfo.(model.User); ok {
					if value.ID == user.ID {
						fmt.Println("-------------------踢除重复登录SessionID:",onlineSessionID)
						middleware.SMgr.EndSessionBy(onlineSessionID)
					}
				}
			}
		}

		//设置变量值
		middleware.SMgr.SetSessionVal(sessionID, "UserInfo", loginUserInfo)
		tokenString := middleware.GenerateToken(&user)
		token = tokenString
	}

	resultBean := model.NewResultBean(false, "用户名或密码错误")
	if exist {
		resultBean = model.NewResultBean(user)
		resultBean.Token = token
	}
	return resultBean
}

func (c *LoginController) PostLoginout() {
	c.Ctx.RemoveCookie("token")
	//如果要设置自定义路径：
	// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))
}

func (c *LoginController) PostUpdatepassword() *model.ResultBean  {
	return model.NewResultBean("")
}

func (c *LoginController) PostGetinfo()*model.ResultBean {
	var user model.User
	c.Ctx.ReadJSON(&user)
	user = c.Service.GetInfo(user.Username)
	resultBean := model.NewResultBean(false,"获取失败")
	if user.ID != 0{
		resultBean = model.NewResultBean(user)
	}
	return resultBean
}



