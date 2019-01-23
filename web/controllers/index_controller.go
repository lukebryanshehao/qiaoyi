package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx iris.Context
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Get() (mvc.Result)  {
	return mvc.View{
		Name: "index-2.html",
		Data: nil,
	}
}
func (c *IndexController) GetWelcome() (mvc.Result)  {
	return mvc.View{
		Name: "welcome.html",
		Data: nil,
	}
}

