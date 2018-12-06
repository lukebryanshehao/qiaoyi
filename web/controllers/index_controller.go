package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {

}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Get(context iris.Context) (mvc.Result)  {
	return mvc.View{
		Name: "index-2.html",
		Data: nil,
	}
}
func (c *IndexController) GetWelcome(context iris.Context) (mvc.Result)  {
	return mvc.View{
		Name: "welcome.html",
		Data: nil,
	}
}

