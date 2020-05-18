// Package routers ...
// @APIVersion 1.0.0
// @Title Shop API
// @Description documents for Shop API
package routers

import (
	"shop-api/controllers"
	"shop-api/services"

	"github.com/astaxie/beego"
)

func init() {
	productService := services.NewProductService()
	categoryService := services.NewCategoryService()
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{
					CategoryService: categoryService,
				},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{
					ProductService: productService,
				},
			),
		),
		beego.NSNamespace("/error",
			beego.NSInclude(
				&controllers.ErrorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
