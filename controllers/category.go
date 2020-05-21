package controllers

import (
	"encoding/json"
	"shop-api/helper"
	"shop-api/services"
	"shop-api/types"
	"shop-api/utils"
	"strconv"

	"github.com/astaxie/beego"
)

// CategoryController operations for Category
type CategoryController struct {
	beego.Controller
	CategoryService services.Category
}

// URLMapping Category controller
func (c *CategoryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post adds a new category
// @Title Post
// @Description create Category
// @Param	body		body 	types.InputAddCategory	true		"body for Category content"
// @Success 201 {int}
// @Failure 403 {int} body is empty
// @Failure 400 {int} Bad Request
// @router / [post]
func (c *CategoryController) Post() {
	var v types.InputAddCategory
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	ormer := helper.NewOrm(true)
	responseCode, id, err := c.CategoryService.Add(ormer, &v)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

type response struct {
	Message string `json:"message"`
}

// GetOne return the category by ID
// @Title Get One
// @Description get Category by ID
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputCategory
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	ormer := helper.NewOrm(false)
	responseCode, result, err := c.CategoryService.GetByID(ormer, id)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetAll retrieves all Category matches certain condition
// @Title Get All
// @Description get Category
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} types.OutputCategory
// @Failure 403 {string} string
// @Header 403 {string} string
// @router / [get]
func (c *CategoryController) GetAll() {
	var limit int64 = 10
	var offset int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	queryString := c.GetString("query")
	query, err := utils.TransformQueryGetAll(queryString)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(types.ResponseCode["Forbidden"])
		c.ServeJSON()
		return
	}
	// sortby: col1,col2
	sortbyString := c.GetString("sortby")
	// order: desc,asc
	orderString := c.GetString("order")
	// query: k:v,k:v
	order, err := utils.TransFormSortFieldOrderGetAll(sortbyString, orderString)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(types.ResponseCode["Forbidden"])
		c.ServeJSON()
		return
	}
	ormer := helper.NewOrm(false)
	responseCode, results, err := c.CategoryService.GetAll(ormer, query, order, offset, limit)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = results
	}
	c.ServeJSON()
}

// Put update category by ID
// @Title Put
// @Description update the Category
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateCategory	true		"body for Category content"
// @Success 200 {object} models.Category
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CategoryController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var v types.InputUpdateCategory
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	ormer := helper.NewOrm(false)
	responseCode, err := c.CategoryService.UpdateByID(ormer, id, &v)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete category by ID
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	ormer := helper.NewOrm(false)
	responseCode, err := c.CategoryService.Delete(ormer, id)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
