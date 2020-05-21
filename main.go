package main

import (
	"shop-api/config"
	"shop-api/controllers"
	_ "shop-api/middleware"
	_ "shop-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	env, _ := config.FromEnv()
	orm.RegisterDataBase("default", env.Driver, env.Sqlconn)
	beego.ErrorController(&controllers.ErrorController{})
	if beego.BConfig.RunMode == "dev" { // Provides swagger and test coverage for dev environment
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.StaticDir["/coverage"] = "coverage"
	}
	beego.Run()
}
