package controllers

import (
	"qiaoyi_back/model"
	"qiaoyi_back/services"
	"qiaoyi_back/repositorys"
	"encoding/json"
	"github.com/kataras/iris"
)

type AreaController struct {
	Service services.AreaService
}

func NewAreaController() *AreaController {
	return &AreaController{Service: services.NewAreaService(repositorys.NewAreaRepository())}
}

func (c *AreaController) PostInsert(context iris.Context) (model.ResultBean)  {
	area := &model.Area{}
	context.ReadJSON(&area)
	id := c.Service.Insert(area)
	resultBean := model.CreateResultWithMsg("添加失败!")
	if id != 0 {
		resultBean = model.CreateResultWithData(nil)
	}
	return resultBean
}

func (c *AreaController) PostDeleteBy(id uint) (model.ResultBean)  {
	flag := c.Service.DeleteByID(id)
	resultBean := model.CreateResultWithMsg("删除失败!")
	if flag {
		resultBean = model.CreateResultWithData(nil)
	}
	return resultBean
}

func (c *AreaController) PostUpdate(context iris.Context) (model.ResultBean)  {
	area := &model.Area{}
	context.ReadJSON(&area)
	flag := c.Service.Update(area)
	resultBean := model.CreateResultWithMsg("修改失败!")
	if flag {
		resultBean = model.CreateResultWithData(nil)
	}
	return resultBean
}

func (c *AreaController) PostAll(context iris.Context) (model.ResultBean)  {
	page := &model.Page{}
	context.ReadJSON(&page)
	var areas = c.Service.PageQuery(page)
	res, err := json.Marshal(areas)
	resultBean := model.CreateResultWithData(res)
	if err != nil {
		resultBean = model.CreateResultWithMsg("")
		panic(err)
	}
	return resultBean
}

func (c *AreaController) GetBy(id int) (model.ResultBean)  {
	area,flag := c.Service.GetByID(id)
	resultBean := model.CreateResultWithMsg("获取失败!")
	if flag {
		resultBean = model.CreateResultWithData(area)
	}
	return resultBean
}

