package controllers

import (
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
	ProductService services.ProductService
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
// @Failure 403 body is empty
// @router / [post]
func (c *ProductController) Post() {
	var v types.InputAddProduct
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode, id, err := c.ProductService.Add(v)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne return the product by ID
// @Title Get One
// @Description get Product by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.OutputProduct
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	responseCode, result, err := c.ProductService.GetByID(id)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = result
	}
	c.ServeJSON()
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
// @Failure 403
// @router / [get]
func (c *ProductController) GetAll() {
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
	responseCode, results, err := c.ProductService.GetAll(query, order, offset, limit)
	c.Ctx.Output.SetStatus(responseCode)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = results
	}
	c.ServeJSON()
}

// Put update product by ID
// @Title Put
// @Description update the Product
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	types.InputUpdateProduct	true		"body for Product content"
// @Success 200 {object} models.Product
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var input types.InputUpdateProduct
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	responseCode, err := c.ProductService.UpdateByID(id, &input)
	c.Ctx.Output.SetStatus(responseCode)
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete product by ID
// @Title Delete
// @Description delete the Product
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	responseCdoe, err := c.ProductService.Delete(id)
	c.Ctx.Output.SetStatus(responseCdoe)
	if err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
