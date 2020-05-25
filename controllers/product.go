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

// ProductController operations for Product
type ProductController struct {
	beego.Controller
	ProductService services.Product
}

// URLMapping ...
func (c *ProductController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post adds a new product to store
// @Title Post
// @Description create Product
// @Param	body		body 	types.InputAddProduct	true		"body for Product content"
// @Success 201 {int}
// @Failure 400 {message: "string"}
// @router / [post]
// @Security apiKey
func (c *ProductController) Post() {
	var input types.InputAddProduct                 // Declare variable type input add product
	json.Unmarshal(c.Ctx.Input.RequestBody, &input) // Parses the JSON-encoded data and input struct
	ctx := context.Background()                     // Declare a new ctx
	id, err := c.ProductService.Add(ctx, input)     // Call service method Add
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}
}

// GetOne return the product by ID
// @Title Get One
// @Description get Product by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputProduct
// @Failure 400 {message: "string"}
// @router /:id [get]
// @Security apiKey
func (c *ProductController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")                // Declare idStr and set equal param id
	id, _ := strconv.ParseInt(idStr, 0, 64)          // Convert idStr string to id int64
	ctx := context.Background()                      // Declare a new context
	result, err := c.ProductService.GetByID(ctx, id) // Call method GetByID
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = result
		c.ServeJSON()
	}
}

// GetAll retrieves all product matches certain condition
// @Title Get All
// @Description get Product
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} types.OutputProduct
// @Failure 400 {message: "string"}
// @router / [get]
// @Security apiKey
func (c *ProductController) GetAll() {
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
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
		return
	}
	sortbyString := c.GetString("sortby") // sortby: col1,col2
	orderString := c.GetString("order")   // order: desc,asc
	order, err := utils.TransFormSortFieldOrderGetAll(sortbyString, orderString)
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
		return
	}
	ctx := context.Background()                                               // Declare a new context
	results, err := c.ProductService.GetAll(ctx, query, order, offset, limit) // Call method GetAll
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = results
		c.ServeJSON()
	}
}

// Put update product by ID
// @Title Put
// @Description update the Product
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateProduct	true		"body for Product content"
// @Success 200 {"OK"}
// @Failure 400 {message: "string"}
// @router /:id [put]
// @Security apiKey
func (c *ProductController) Put() {
	idStr := c.Ctx.Input.Param(":id")                   // Declare idStr and set it equal param id
	id, _ := strconv.ParseInt(idStr, 0, 64)             // Declare id and convert idStr to id
	var input types.InputUpdateProduct                  // Declare input type input update product
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)     // Parses the JSON-encoded data and input struct
	ctx := context.Background()                         // Declare a context
	err := c.ProductService.UpdateByID(ctx, id, &input) // Call method UpdateByID
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}

// Delete product by ID
// @Title Delete
// @Description delete the Product
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 "OK"
// @Failure 400 {message: "string"}
// @router /:id [delete]
// @Security apiKey
func (c *ProductController) Delete() {
	idStr := c.Ctx.Input.Param(":id")       // Declare idStr and set equal id
	id, _ := strconv.ParseInt(idStr, 0, 64) // Declare id and convert idStr to id
	ctx := context.Background()             // Declare a context
	err := c.ProductService.Delete(ctx, id) // Call method Delete
	if err != nil {
		c.Ctx.Input.SetParam("errMessage", err.Error())
	} else {
		c.Data["json"] = "OK"
		c.ServeJSON()
	}
}
