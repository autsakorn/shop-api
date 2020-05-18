package services

import (
	"fmt"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"
	"time"
)

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
func (ps ProductService) Add(product types.InputAddProduct) (id int64, err error) {
	v := models.Product{
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
	id, err = ps.Storage.Product.Add(&v)
	return
}

// Delete ...
func (ps ProductService) Delete(id int64) (err error) {
	modelProduct := models.Product{
		ID: id,
	}
	var num int64
	num, err = ps.Storage.Product.Delete(&modelProduct)
	fmt.Println("num", num)
	return
}

// GetByID ...
func (ps ProductService) GetByID(id int64) (product models.Product, err error) {
	fmt.Println("id", id)
	product, err = ps.Storage.Product.GetByID(id)
	return
}

// GetAll ...
func (ps ProductService) GetAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	ml, err = ps.Storage.Product.GetAll(query, fields, sortby, order, offset, limit)
	return
}

// UpdateByID ...
func (ps ProductService) UpdateByID(id int64, product *types.InputUpdateProduct) (err error) {
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
	_, err = ps.Storage.Product.UpdateByID(&m)
	return
}
