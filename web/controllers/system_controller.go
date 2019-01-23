package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
	"strconv"
)

type SystemController struct {
	Service services.SystemService
	Ctx iris.Context
}

func NewSystemController() *SystemController {
	return &SystemController{Service: services.NewSystemService(repositorys.NewSystemRepository())}
}

func (c *SystemController) Get() (mvc.Result)  {
	page := &model.Page{}
	pageSize,err1 := strconv.Atoi(c.Ctx.Request().FormValue("PageSize"))
	pageIndex,err2 := strconv.Atoi(c.Ctx.Request().FormValue("PageIndex"))
	if err1 != nil || err2 != nil {
		//panic(err1)
		//panic(err2)
	}
	page.PageSize = pageSize
	page.PageIndex = pageIndex
	allcount,settings := c.Service.PageQuery(page)
	resultBean := model.NewResultPage(settings,allcount)
	maps := map[string]interface{}{
		"ResultBean":     resultBean,
	}
	return mvc.View{
		Name: "system-base.html",
		Data: maps,
	}
}
