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

// BrandController operations for Brand
type BrandController struct {
	beego.Controller
	BrandService services.BrandService
}

// URLMapping ...
func (c *BrandController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Brand
// @Param	body		body 	models.Brand	true		"body for Brand content"
// @Success 201 {int} ID
// @Failure 400 {message: "string"}
// @router / [post]
func (c *BrandController) Post() {
	var v types.InputAddBrand                   // Declare a variable input add category
	json.Unmarshal(c.Ctx.Input.RequestBody, &v) // Parses the JSON-encoded data and input struct
	ctx := context.Background()                 // Declare a context
	id, err := c.BrandService.Add(ctx, &v)      // Call service method Add
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}

}

// GetOne ...
// @Title Get One
// @Description get Brand by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputBrand
// @Failure 400 {message: "string"}
// @router /:id [get]
func (c *BrandController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")              // Get id from param
	id, _ := strconv.ParseInt(idStr, 0, 64)        // Convert id(string) to int64
	ctx := context.Background()                    // Create a context
	result, err := c.BrandService.GetByID(ctx, id) // Call service method GetByID
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// GetAll ...
// @Title Get All
// @Description get Brand
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} types.OutputBrand
// @Failure 400 {message: "string"}
// @router / [get]
func (c *BrandController) GetAll() {
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
	ctx := context.Background()                                             // Create a context
	results, err := c.BrandService.GetAll(ctx, query, order, offset, limit) // Cal service method GetAll
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = results
		c.ServeJSON()
	}
}

// Put ...
// @Title Put
// @Description update the Brand
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateBrand	true		"body for Brand content"
// @Success 200 OK
// @Failure 400 {message: "string"}
// @router /:id [put]
func (c *BrandController) Put() {
	idStr := c.Ctx.Input.Param(":id")                 // Get id from param and declare a idStr variable
	id, _ := strconv.ParseInt(idStr, 0, 64)           // Convert idStr to id type int64
	var input types.InputUpdateBrand                  // Declare input type InputUpdateCategory
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)   // Parses the JSON-encoded data and input struct
	ctx := context.Background()                       // Declare context
	err := c.BrandService.UpdateByID(ctx, id, &input) // Call UpdateByID method
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}

// Delete ...
// @Title Delete
// @Description delete the Brand
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 OK
// @Failure 400 {message: "string"}
// @router /:id [delete]
func (c *BrandController) Delete() {
	idStr := c.Ctx.Input.Param(":id")       // Get id from param and declare a idStr variable
	id, _ := strconv.ParseInt(idStr, 0, 64) // Convert idStr to id type int64
	ctx := context.Background()             // Declare context
	err := c.BrandService.Delete(ctx, id)   // Call Delete method
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}
