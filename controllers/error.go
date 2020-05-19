package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController ...
type ErrorController struct {
	beego.Controller
}

// Error404 ...
func (c *ErrorController) Error404() {
	c.Data["content"] = "Page not found"
	c.TplName = "error.tpl"
	c.Render()
}

// Error401 ...
func (c *ErrorController) Error401() {
	c.Data["content"] = "Unauthorized"
	c.TplName = "error.tpl"
	c.Render()
}

// Error500 ...
func (c *ErrorController) Error500() {
	c.Data["content"] = "Internal server error"
	c.TplName = "error.tpl"
	c.Render()
}

// ErrorDb ...
func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error.tpl"
	c.Render()
}
