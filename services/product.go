package services

import (
	"errors"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"

	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/copier"
)

// Product represents all possible actions available for product services
type Product interface {
	Add(ormer orm.Ormer, product types.InputAddProduct) (responseCode int, id int64, err error)
	Delete(ormer orm.Ormer, id int64) (responseCode int, err error)
	GetByID(ormer orm.Ormer, id int64) (responseCode int, product types.OutputProduct, err error)
	GetAll(
		ormer orm.Ormer,
		query map[string]string,
		order []string,
		offset int64,
		limit int64,
	) (responseCode int, results []types.OutputProduct, err error)
	UpdateByID(ormer orm.Ormer, id int64, product *types.InputUpdateProduct) (responseCode int, err error)
}

// ProductService defines propertie
type ProductService struct {
	Storage storage.Storage
}

// NewProductService map storage and return ProductService
func NewProductService() (ps ProductService) {
	ps.Storage = storage.NewStorage()
	return
}

// Add service for add a new product
func (ps ProductService) Add(ormer orm.Ormer, product types.InputAddProduct) (responseCode int, id int64, err error) {
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
	id, err = ps.Storage.Product.Add(ormer, &inputModel)
	if id > 0 {
		responseCode = types.ResponseCode["CreatedSuccess"]
		return
	}
	return
}

// Delete service for delete product by ID
func (ps ProductService) Delete(ormer orm.Ormer, id int64) (responseCode int, err error) {
	responseCode = types.ResponseCode["Success"]
	modelProduct := models.Product{
		ID: id,
	}
	num, err := ps.Storage.Product.Delete(ormer, &modelProduct)
	if num < 1 {
		errorMessage := "Not Found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetByID service retrieve product by ID
func (ps ProductService) GetByID(ormer orm.Ormer, id int64) (responseCode int, result types.OutputProduct, err error) {
	responseCode = types.ResponseCode["Success"]
	product, err := ps.Storage.Product.GetByID(ormer, id)
	copier.Copy(&result, &product)
	if err != nil {
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetAll service for retrieves all product matches certain condition
func (ps ProductService) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (responseCode int, results []types.OutputProduct, err error) {
	products, err := ps.Storage.Product.GetAll(ormer, query, order, offset, limit)
	copier.Copy(&results, &products)
	responseCode = types.ResponseCode["Success"]
	return
}

// UpdateByID service for update product by ID
func (ps ProductService) UpdateByID(ormer orm.Ormer, id int64, product *types.InputUpdateProduct) (responseCode int, err error) {
	responseCode = types.ResponseCode["Success"]
	dataProduct, err := ps.Storage.Product.GetByID(ormer, id)
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
	num, err := ps.Storage.Product.UpdateByID(ormer, &m)
	if num < 1 {
		errorMessage := "Not Found"
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
		return
	}
	return
}
