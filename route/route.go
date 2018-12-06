package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/web/controllers"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/area")).Handle(controllers.NewAreaController())
	mvc.New(app.Party("/index")).Handle(controllers.NewIndexController())
	mvc.New(app.Party("/login")).Handle(controllers.NewLoginController())
	//mvc.Configure(app.Party("/area"), area)

}

//func area(app *mvc.Application) {
//	repo := repositorys.NewAreaRepository()
//	areaService := services.NewAreaService(repo)
//	app.Register(areaService)
//
//	app.Handle(new(controllers.AreaController))
//}