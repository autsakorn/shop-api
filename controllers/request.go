package controllers

import (
	"context"
	"encoding/json"
	"shop-api/services"
	"shop-api/types"
	"shop-api/utils"
	"strconv"

	"github.com/astaxie/beego"
)

// RequestController operations for Request
type RequestController struct {
	beego.Controller
	RequestService services.RequestService
}

// URLMapping ...
func (c *RequestController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Request
// @Param	body		body 	types.InputAddRequest	true		"body for Request content"
// @Success 201 {int} types.OutputRequest
// @Failure 403 body is empty
// @router / [post]
// @Security apiKey
func (c *RequestController) Post() {
	var v types.InputAddRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v) // Parses the JSON-encoded data and input struct
	ctx := context.Background()                 // Declare a context
	id, err := c.RequestService.Add(ctx, &v)    // Call service method Add
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}
}

// GetOne ...
// @Title Get One
// @Description get Request by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputAddRequest
// @Failure 403 :id is empty
// @router /:id [get]
// @Security apiKey
func (c *RequestController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")                // Get id from param
	id, _ := strconv.ParseInt(idStr, 0, 64)          // Convert id(string) to int64
	ctx := context.Background()                      // Create a context
	result, err := c.RequestService.GetByID(ctx, id) // Call service method GetByID
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// GetAll ...
// @Title Get All
// @Description get Request
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	request	query	string	false	"Request corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Request
// @Failure 403
// @router / [get]
// @Security apiKey
func (c *RequestController) GetAll() {
	var limit int64 = 10
	var offset int64
	if v, err := c.GetInt64("limit"); err == nil { // limit: 10 (default is 10)
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil { // offset: 0 (default is 0)
		offset = v
	}
	queryString := c.GetString("query") // query: k:v,k:v
	query, err := utils.TransformQueryGetAll(queryString)
	if err != nil { // Handle invalid form
		c.Ctx.Input.SetParam("errMessage", err.Error())
		return
	}
	sortbyString := c.GetString("sortby")   // sortby: col1,col2
	requestString := c.GetString("request") // request: desc,asc
	request, err := utils.TransFormSortFieldOrderGetAll(sortbyString, requestString)
	if err != nil { // Handle invalid form
		c.Ctx.Input.SetParam("errMessage", err.Error())
		return
	}
	ctx := context.Background()                                                 // Create a context
	results, err := c.RequestService.GetAll(ctx, query, request, offset, limit) // Cal service method GetAll
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = results
		c.ServeJSON()
	}
}

// Put ...
// @Title Put
// @Description update the Request
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateRequest      	true		"body for Request content"
// @Success 200 {object} types.OutputRequest
// @Failure 403 :id is not int
// @router /:id [put]
// @Security apiKey
func (c *RequestController) Put() {
	idStr := c.Ctx.Input.Param(":id")                   // Get id from param and declare a idStr variable
	id, _ := strconv.ParseInt(idStr, 0, 64)             // Convert idStr to id type int64
	var input types.InputUpdateRequest                  // Declare input type InputUpdateRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)     // Parses the JSON-encoded data and input struct
	ctx := context.Background()                         // Declare context
	err := c.RequestService.UpdateByID(ctx, id, &input) // Call UpdateByID method
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}
