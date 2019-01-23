package controllers

import (
	"github.com/kataras/iris"
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
	"qiaoyi_back/services"
)

type AreaController struct {
	Service services.AreaService
	Ctx iris.Context
}

func NewAreaController() *AreaController {
	return &AreaController{Service: services.NewAreaService(repositorys.NewAreaRepository())}
}

func (c *AreaController) PostInsert() (*model.ResultBean)  {
	area := &model.Area{}
	c.Ctx.ReadJSON(&area)
	id := c.Service.Insert(area)
	resultBean := model.NewResultBean(false,"添加失败!")
	if id != 0 {
		resultBean = model.NewResultBean(nil)
	}
	return resultBean
}

func (c *AreaController) PostDeleteBy(id uint) (*model.ResultBean)  {
	flag := c.Service.DeleteByID(id)
	resultBean := model.NewResultBean(false,"删除失败!")
	if flag {
		resultBean = model.NewResultBean(nil)
	}
	return resultBean
}

func (c *AreaController) PostUpdate() (*model.ResultBean)  {
	area := &model.Area{}
	c.Ctx.ReadJSON(&area)
	flag := c.Service.Update(area)
	resultBean := model.NewResultBean(false,"修改失败!")
	if flag {
		resultBean = model.NewResultBean(nil)
	}
	return resultBean
}

func (c *AreaController) PostAll() (*model.ResultPage)  {
	page := &model.Page{}
	c.Ctx.ReadJSON(&page)
	allcount,areas := c.Service.PageQuery(page)
	//res, err := json.Marshal(areas)
	resultBean := model.NewResultPage(areas,allcount)
	if allcount < 1 {
		resultBean = model.NewResultPage("没有获取到数据!",0)
	}
	return resultBean
}

func (c *AreaController) GetBy(id int) (*model.ResultBean)  {
	area,flag := c.Service.GetByID(id)
	resultBean := model.NewResultBean("获取失败!")
	if flag {
		resultBean = model.NewResultBean(area)
	}
	return resultBean
}

func (c *AreaController) GetTree() (*model.ResultBean)  {
	id := c.Ctx.Request().FormValue("id")
	areaarr := c.Service.GetTree(id)
	resultBean := model.NewResultBean("获取失败!")
	if len(areaarr) > 0 {
		resultBean = model.NewResultBean(areaarr)
	}
	return resultBean
}


