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
	Ctx iris.Context
}

func NewUserController() *UserController {
	return &UserController{Service: services.NewUserService(repositorys.NewUserRepository())}
}

func (c *UserController) PostInsert() (*model.ResultBean)  {
	user := &model.User{}
	c.Ctx.ReadJSON(&user)
	flag,user := c.Service.Save(user)
	resultBean := model.NewResultBean(false,"添加失败!")
	if flag {
		resultBean = model.NewResultBean("添加成功!")
	}
	return resultBean
}

func (c *UserController) PostDeleteBy(id uint) (*model.ResultBean)  {
	flag := c.Service.DeleteByID(id)
	resultBean := model.NewResultBean(false,"删除失败!")
	if flag {
		resultBean = model.NewResultBean(nil)
	}
	return resultBean
}

func (c *UserController) PostUpdate() (*model.ResultBean)  {
	user := &model.User{}
	c.Ctx.ReadJSON(&user)
	flag,user := c.Service.Save(user)
	resultBean := model.NewResultBean(false,"修改失败!")
	if flag {
		resultBean = model.NewResultBean("修改成功!")
	}
	return resultBean
}

func (c *UserController) PostAll() (*model.ResultPage)  {
	page := &model.Page{}
	c.Ctx.ReadJSON(&page)
	allcount,users := c.Service.PageQuery(page)
	_, err := json.Marshal(users)
	resultBean := model.NewResultPage(users,allcount)
	if err != nil {
		resultBean = model.NewResultPage("",0)
		panic(err)
	}
	return resultBean
}

func (c *UserController) GetBy(id uint) (*model.ResultBean)  {
	flag,user := c.Service.GetByID(id)
	resultBean := model.NewResultBean(false,"获取失败!")
	if flag {
		resultBean = model.NewResultBean(user)
	}
	return resultBean
}

