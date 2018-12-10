package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
	"qiaoyi_back/utils"
	"strconv"
)

type RoleController struct {
	Service services.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{Service: services.NewRoleService(repositorys.NewRoleRepository())}
}

func (c *RoleController) Get(context iris.Context) (mvc.Result)  {
	page := &model.Page{}
	pageSize,err1 := strconv.Atoi(context.Request().FormValue("PageSize"))
	pageIndex,err2 := strconv.Atoi(context.Request().FormValue("PageIndex"))
	if err1 != nil || err2 != nil {
		//panic(err1)
		//panic(err2)
	}
	page.PageSize = pageSize
	page.PageIndex = pageIndex
	//context.ReadJSON(&page)
	allcount,roles := c.Service.PageQuery(page)
	resultBean := model.CreateResultWithCountAndData(allcount,roles)
	//json, err := json.Marshal(resultBean)
	//if err != nil {
	//	panic(err)
	//}
	var po utils.PageOptions                                                //定义一个分页对象
	po.EnableFirstLastLink = true                                           //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                                              //是否显示上一页下一页 默认为false
	po.Currentpage = pageIndex                                              //传递当前页数,默认为1
	po.PageSize = pageSize                                                  //页面大小  默认为20
	_, pagerhtml := utils.GetPagerLinks(allcount,&po)						 //返回总页数,html
	maps := map[string]interface{}{
		"ResultBean":     resultBean,
		"totalItem":     allcount,
		"pagerhtml":     pagerhtml,
	}
	return mvc.View{
		Name: "admin-role.html",
		Data: maps,
	}
}

func (c *RoleController) GetBy(id int) (mvc.Result)  {
	role,flag := c.Service.GetByID(id)

	resultBean := model.CreateResultWithMsg("获取失败!")
	if flag {
		resultBean = model.CreateResultWithData(role)
	}
	maps := map[string]interface{}{
		"ResultBean":     resultBean,
	}
	return mvc.View{
		Name: "admin-role-add.html",
		Data: maps,
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

func (c *RoleController) PostUpdate(context iris.Context)  (model.ResultBean) {
	role := &model.Role{}
	rid := context.Request().FormValue("ID")
	if rid != ""{
		id,_ := strconv.Atoi(rid)
		role.ID = uint(id)
		role2,flag := c.Service.GetByID(id)
		if flag {
			role.CreatedAt = role2.CreatedAt
		}
	}

	rolename := context.Request().FormValue("Rolename")
	weight,_ := strconv.Atoi(context.Request().FormValue("Weight"))
	remark := context.Request().FormValue("Remark")

	role.Rolename = rolename
	role.Weight = uint(weight)
	role.Remark = remark
	flag,role := c.Service.Save(role)
	msg := "添加"
	if rid != "" {
		msg = "修改"
	}
	resultBean := model.CreateResultWithMsg(msg+"失败!")
	if flag {
		resultBean = model.CreateResultWithData(msg+"成功!")
	}
	return resultBean
}
