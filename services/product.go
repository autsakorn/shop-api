package services

import (
	"errors"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"

	"github.com/jinzhu/copier"
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
		errorMessage := "Not Found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetByID ...
func (ps ProductService) GetByID(id int64) (responseCode int, result types.OutputProduct, err error) {
	responseCode = types.ResponseCode["Success"]
	product, err := ps.Storage.Product.GetByID(id)
	copier.Copy(&result, &product)
	if err != nil {
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetAll ...
func (ps ProductService) GetAll(
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (responseCode int, results []types.OutputProduct, err error) {
	products, err := ps.Storage.Product.GetAll(query, order, offset, limit)
	copier.Copy(&results, &products)
	responseCode = types.ResponseCode["Success"]
	return
}

// UpdateByID ...
func (ps ProductService) UpdateByID(id int64, product *types.InputUpdateProduct) (responseCode int, err error) {
	responseCode = types.ResponseCode["Success"]
	dataProduct, err := ps.Storage.Product.GetByID(id)
	m := models.Product{
		ID:        dataProduct.ID,
		Name:      product.Name,
		Detail:    product.Detail,
		Brand:     product.Brand,
		Model:     product.Model,
		Cost:      product.Cost,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: dataProduct.CreatedAt,
		Category: &models.Category{
			ID: product.Category.ID,
		},
	}
	num, err := ps.Storage.Product.UpdateByID(&m)
	if num < 1 {
		errorMessage := "Not Found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
		return
	}
	return
}
