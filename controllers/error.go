package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController handle abort error
type ErrorController struct {
	beego.Controller
}

// Error404 return Page not found if abort 400
func (c *ErrorController) Error404() {
	c.Data["content"] = "Page not found"
	c.TplName = "error.tpl"
	c.Render()
}

// Error401 return Unauthorized if abort 401
func (c *ErrorController) Error401() {
	c.Data["content"] = "Unauthorized"
	c.TplName = "error.tpl"
	c.Render()
}

// Error500 return Internal server error if abort 500
func (c *ErrorController) Error500() {
	c.Data["content"] = "Internal server error"
	c.TplName = "error.tpl"
	c.Render()
}

// ErrorDb return database is now down if about Db
func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error.tpl"
	c.Render()
}
