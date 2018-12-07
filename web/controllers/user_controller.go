package controllers

import (
	"qiaoyi_back/model"
	"qiaoyi_back/services"
	"qiaoyi_back/repositorys"
	"encoding/json"
	"github.com/kataras/iris"
)

type UserController struct {
	Service services.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: services.NewUserService(repositorys.NewUserRepository())}
}

func (c *UserController) PostInsert(context iris.Context) (model.ResultBean)  {
	user := &model.User{}
	context.ReadJSON(&user)
	flag,user := c.Service.Save(user)
	resultBean := model.CreateResultWithMsg("添加失败!")
	if flag {
		resultBean = model.CreateResultWithData("添加成功!")
	}
	return resultBean
}

func (c *UserController) PostDeleteBy(id uint) (model.ResultBean)  {
	flag := c.Service.DeleteByID(id)
	resultBean := model.CreateResultWithMsg("删除失败!")
	if flag {
		resultBean = model.CreateResultWithData(nil)
	}
	return resultBean
}

func (c *UserController) PostUpdate(context iris.Context) (model.ResultBean)  {
	user := &model.User{}
	context.ReadJSON(&user)
	flag,user := c.Service.Save(user)
	resultBean := model.CreateResultWithMsg("修改失败!")
	if flag {
		resultBean = model.CreateResultWithData("修改成功!")
	}
	return resultBean
}

func (c *UserController) PostAll(context iris.Context) (model.ResultBean)  {
	page := &model.Page{}
	context.ReadJSON(&page)
	allcount,users := c.Service.PageQuery(page)
	_, err := json.Marshal(users)
	resultBean := model.CreateResultWithCountAndData(allcount,users)
	if err != nil {
		resultBean = model.CreateResultWithMsg("")
		panic(err)
	}
	return resultBean
}

func (c *UserController) GetBy(id uint) (model.ResultBean)  {
	flag,user := c.Service.GetByID(id)
	resultBean := model.CreateResultWithMsg("获取失败!")
	if flag {
		resultBean = model.CreateResultWithData(user)
	}
	return resultBean
}

