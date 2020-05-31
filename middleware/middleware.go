package middleware

import (
	ctxContext "context"
	"shop-api/services"
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
	AllowHeaders:     []string{"X-API-KEY", "Content-Type"},
}

var auth = func(ctx *context.Context) {
	xAPIKey := ctx.Request.Header.Get("X-API-KEY")
	clientService := services.NewClientService()
	ctxBackground := ctxContext.Background()
	err := clientService.VerifyXApiKey(ctxBackground, xAPIKey)
	if err != nil {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(types.ErrorFormat{Message: "Invalid X-API-KEY"}, false, false)
		return
	}
}

func init() {
	beego.InsertFilter("*", beego.BeforeStatic, cors.Allow(corsOption), true)
	beego.InsertFilter("/v1/product/*", beego.BeforeRouter, auth, false)
	beego.InsertFilter("/v1/request/*", beego.BeforeRouter, auth, false)
	beego.InsertFilter("*", beego.AfterExec, afterExecMiddleware, false)
}
