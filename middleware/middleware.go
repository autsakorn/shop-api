package middleware

import (
	"shop-api/types"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

var afterExecMiddleware = func(ctx *context.Context) {
	errMessage := ctx.Input.Param("errMessage")
	if errMessage != "" {
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(types.ErrorFormat{Message: errMessage}, false, false)
	}
}

var corsOption = &cors.Options{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowCredentials: true,
}

func init() {
	beego.InsertFilter("*", beego.BeforeStatic, cors.Allow(corsOption), true)
	beego.InsertFilter("*", beego.AfterExec, afterExecMiddleware, false)
}
