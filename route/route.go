package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"qiaoyi_back/web/controllers"
	"qiaoyi_back/web/middleware"
)

func InitRouter(app *iris.Application) {
	go app.Use(middleware.SaveSystemLog)
	mvc.New(app.Party("/login")).Handle(controllers.NewLoginController())

	app.Use(middleware.CheckSession)
	app.Use(middleware.GetJWT().Serve)
	mvc.New(app.Party("/area")).Handle(controllers.NewAreaController())
	mvc.New(app.Party("/index")).Handle(controllers.NewIndexController())
	mvc.New(app.Party("/role")).Handle(controllers.NewRoleController())
	mvc.New(app.Party("/user")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/system")).Handle(controllers.NewSystemController())
	mvc.New(app.Party("/file")).Handle(controllers.NewFileController())
	//mvc.Configure(app.Party("/area"), area)

}

//func area(app *mvc.Application) {
//	repo := repositorys.NewAreaRepository()
//	areaService := services.NewAreaService(repo)
//	app.Register(areaService)
//
//	app.Handle(new(controllers.AreaController))
//}