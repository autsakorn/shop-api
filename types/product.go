package types

import (
	"shop-api/models"
)

// Response ...
type Response struct {
	Message string `json:"message"`
}

type productCategory struct {
	ID int64 `json:"id"`
}

// InputAddProduct defines properties of input new product
type InputAddProduct struct {
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Detail   string          `json:"detail"`
	Brand    string          `json:"brand"`
	Model    string          `json:"model"`
	Quantity int             `json:"quantity"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// InputDeleteProduct defines properties of input delete product
type InputDeleteProduct struct {
	ID int64 `json:"id"`
}

// InputUpdateProduct defines properties of input update and upsert product
type InputUpdateProduct struct {
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Detail   string          `json:"detail"`
	Brand    string          `json:"brand"`
	Model    string          `json:"model"`
	Quantity int             `json:"quantity"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// OutputCreateProduct defines properties of input create product
type OutputCreateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
}

// OutputDeleteProduct defines properties of output delete product
type OutputDeleteProduct struct {
	Deleted bool   `json:"deleted"`
	Message string `json:"message"`
}

// OutputUpdateProduct defines properties of output update product
type OutputUpdateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
	Updated bool   `json:"updated"`
}

// OutputProduct defines properties of output find product
type OutputProduct struct {
	Message string            `json:"message"`
	Data    []*models.Product `json:"data"`
}

// OutputProducts defines properties of output find product
type OutputProducts struct {
	Message string            `json:"message"`
	Totals  int64             `json:"totals"`
	Data    []*models.Product `json:"data"`
}
