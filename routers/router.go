// Package routers ...
// @APIVersion 1.0.0
// @Title Shop API
// @Description documents for Shop API
package routers

import (
	"shop-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
