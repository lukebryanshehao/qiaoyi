package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
)

type RoleController struct {
	Service services.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{Service: services.NewRoleService(repositorys.NewRoleRepository())}
}

func (c *RoleController) Post(context iris.Context) (mvc.Result)  {
	page := &model.Page{}
	context.ReadJSON(&page)
	roles := c.Service.PageQuery(page)
	return mvc.View{
		Name: "admin-role.html",
		Data: model.CreateResultWithData(roles),
	}
}

func (c *RoleController) PostDelete(context iris.Context)  (model.ResultBean) {
	role := &model.Role{}
	context.ReadJSON(&role)
	flag := c.Service.DeleteByID(role.ID)
	resultBean := model.CreateResultWithMsg("删除失败!")
	if flag {
		resultBean = model.CreateResultWithData("删除成功!")
	}
	return resultBean
}

func (c *RoleController) Put(context iris.Context)  (model.ResultBean) {
	role := &model.Role{}
	context.ReadJSON(&role)
	flag,role := c.Service.Save(role)
	msg := "添加"
	if role.ID != 0 {
		msg = "修改"
	}
	resultBean := model.CreateResultWithMsg(msg+"失败!")
	if flag {
		resultBean = model.CreateResultWithData(msg+"成功!")
	}
	return resultBean
}
