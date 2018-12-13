package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"qiaoyi_back/datasource"
	"qiaoyi_back/route"
)

func main() {
	defer datasource.DB.Close()
	datasource.Createtable()//创建数据库
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.StaticWeb("/src", datasource.DBconfig.StaticPath)

	route.InitRouter(app)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	app.Use(iris.Gzip, logger.New(), crs)
	app.AllowMethods(iris.MethodOptions)

	app.Run(
		iris.Addr(":"+datasource.DBconfig.Port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

