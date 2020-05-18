package controllers

import (
	"encoding/json"
	"errors"
	"shop-api/services"
	"shop-api/types"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// CategoryController operations for Category
type CategoryController struct {
	beego.Controller
	CategoryService services.CategoryService
}

// URLMapping ...
func (c *CategoryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Category
// @Param	body		body 	types.InputAddCategory	true		"body for Category content"
// @Success 201 {int} models.Category
// @Failure 403 {int} body is empty
// @Failure 400 {int} Bad Request
// @router / [post]
func (c *CategoryController) Post() {
	var v types.InputAddCategory
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if responseCode, err := c.CategoryService.Add(&v); err == nil {
		c.Ctx.Output.SetStatus(responseCode)
		c.Data["json"] = v
	} else {
		c.Ctx.Output.SetStatus(responseCode)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

type response struct {
	Message string `json:"message"`
}

// GetOne ...
// @Title Get One
// @Description get Category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	responseCode, result, err := c.CategoryService.GetByID(id)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Category
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 403
// @router / [get]
func (c *CategoryController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	responseCode, results, err := c.CategoryService.GetAll(query, fields, sortby, order, offset, limit)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = results
	}
	c.ServeJSON()
}

// Put ...
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
	responseCode, err := c.CategoryService.UpdateByID(id, &v)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	responseCode, err := c.CategoryService.Delete(id)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
