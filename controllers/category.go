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
// @Failure 400 {message: "string"}
// @router / [post]
func (c *CategoryController) Post() {
	var v types.InputAddCategory                // Declare a variable input add category
	json.Unmarshal(c.Ctx.Input.RequestBody, &v) // Parses the JSON-encoded data and input struct
	// ormHelper := helper.NewOrm()                         // Declare a new orm
	ctx := context.Background()               // Declare a context
	id, err := c.CategoryService.Add(ctx, &v) // Call service method Add
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}
}

// GetOne return the category by ID
// @Title Get One
// @Description get Category by ID
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputCategory
// @Failure 400 {message: "string"}
// @router /:id [get]
func (c *CategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")                 // Get id from param
	id, _ := strconv.ParseInt(idStr, 0, 64)           // Convert id(string) to int64
	ctx := context.Background()                       // Create a context
	result, err := c.CategoryService.GetByID(ctx, id) // Call service method GetByID
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = result
		c.ServeJSON()
	}
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
// @Failure 400 {message: "string"}
// @router / [get]
func (c *CategoryController) GetAll() {
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
	sortbyString := c.GetString("sortby") // sortby: col1,col2
	orderString := c.GetString("order")   // order: desc,asc
	order, err := utils.TransFormSortFieldOrderGetAll(sortbyString, orderString)
	if err != nil { // Handle invalid form
		c.Ctx.Input.SetParam("errMessage", err.Error())
		return
	}
	ctx := context.Background()                                                // Create a context
	results, err := c.CategoryService.GetAll(ctx, query, order, offset, limit) // Cal service method GetAll
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = results
		c.ServeJSON()
	}
}

// Put update category by ID
// @Title Put
// @Description update the Category
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateCategory	true		"body for Category content"
// @Success 200 "OK"
// @Failure 400 {message: "string"}
// @router /:id [put]
func (c *CategoryController) Put() {
	idStr := c.Ctx.Input.Param(":id")                    // Get id from param and declare a idStr variable
	id, _ := strconv.ParseInt(idStr, 0, 64)              // Convert idStr to id type int64
	var input types.InputUpdateCategory                  // Declare input type InputUpdateCategory
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)      // Parses the JSON-encoded data and input struct
	ctx := context.Background()                          // Declare context
	err := c.CategoryService.UpdateByID(ctx, id, &input) // Call UpdateByID method
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}

// Delete category by ID
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 "OK"
// @Failure 400 {message: "string"}
// @router /:id [delete]
func (c *CategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")        // Get id from param and declare a idStr variable
	id, _ := strconv.ParseInt(idStr, 0, 64)  // Convert idStr to id type int64
	ctx := context.Background()              // Declare context
	err := c.CategoryService.Delete(ctx, id) // Call Delete method
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}
