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
}

func NewSystemController() *SystemController {
	return &SystemController{Service: services.NewSystemService(repositorys.NewSystemRepository())}
}

func (c *SystemController) Get(context iris.Context) (mvc.Result)  {
	page := &model.Page{}
	pageSize,err1 := strconv.Atoi(context.Request().FormValue("PageSize"))
	pageIndex,err2 := strconv.Atoi(context.Request().FormValue("PageIndex"))
	if err1 != nil || err2 != nil {
		//panic(err1)
		//panic(err2)
	}
	page.PageSize = pageSize
	page.PageIndex = pageIndex
	allcount,settings := c.Service.PageQuery(page)
	resultBean := model.CreateResultWithCountAndData(allcount,settings)
	maps := map[string]interface{}{
		"ResultBean":     resultBean,
	}
	return mvc.View{
		Name: "system-base.html",
		Data: maps,
	}
}
