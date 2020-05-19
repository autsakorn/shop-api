package services

import (
	"errors"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"
	"time"
)

// Product defines action
type Product interface {
	Add(product types.InputAddProduct) (responseCode int, id int64, err error)
	Delete(id int64) (responseCode int, err error)
	GetByID(id int64) (responseCode int, product models.Product, err error)
	GetAll(query map[string]string, fields []string, sortby []string, order []string,
		offset int64, limit int64) (responseCode int, results []interface{}, err error)
	UpdateByID(id int64, product *types.InputUpdateProduct) (responseCode int, err error)
}

// ProductService ...
type ProductService struct {
	Storage storage.Storage
}

// NewProductService ...
func NewProductService() (ps ProductService) {
	ps.Storage = storage.NewStorage()
	return
}

// Add ...
func (ps ProductService) Add(product types.InputAddProduct) (responseCode int, id int64, err error) {
	responseCode = types.ResponseCode["BadRequest"]
	inputModel := models.Product{
		Name:     product.Name,
		Detail:   product.Detail,
		Brand:    product.Brand,
		Model:    product.Model,
		Quantity: product.Quantity,
		Price:    product.Price,
		Cost:     product.Cost,
		Category: &models.Category{
			ID: product.Category.ID,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err = ps.Storage.Product.Add(&inputModel)
	if id > 0 {
		responseCode = types.ResponseCode["CreatedSuccess"]
		return
	}
	return
}

// Delete ...
func (ps ProductService) Delete(id int64) (responseCode int, err error) {
	responseCode = types.ResponseCode["Success"]
	modelProduct := models.Product{
		ID: id,
	}
	num, err := ps.Storage.Product.Delete(&modelProduct)
	if num < 1 {
		errorMessage := "Not found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetByID ...
func (ps ProductService) GetByID(id int64) (responseCode int, product models.Product, err error) {
	product, err = ps.Storage.Product.GetByID(id)
	responseCode = types.ResponseCode["Success"]
	return
}

// GetAll ...
func (ps ProductService) GetAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (responseCode int, results []interface{}, err error) {
	results, err = ps.Storage.Product.GetAll(query, fields, sortby, order, offset, limit)
	responseCode = types.ResponseCode["Success"]
	return
}

// UpdateByID ...
func (ps ProductService) UpdateByID(id int64, product *types.InputUpdateProduct) (responseCode int, err error) {
	responseCode = types.ResponseCode["Success"]
	dataProduct, err := ps.Storage.Product.GetByID(id)
	m := models.Product{
		ID:        id,
		Name:      product.Name,
		Detail:    product.Detail,
		Brand:     product.Brand,
		Model:     product.Model,
		Cost:      product.Cost,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: dataProduct.CreatedAt,
		UpdatedAt: time.Now(),
		Category: &models.Category{
			ID: product.Category.ID,
		},
	}
	num, err := ps.Storage.Product.UpdateByID(&m)
	if num < 1 {
		errorMessage := "Not found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
		return
	}
	return
}
