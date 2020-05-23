package services

import (
	"context"
	"database/sql"
	"errors"
	"shop-api/helper"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"

	"github.com/jinzhu/copier"
)

// Product represents all possible actions available for product services
type Product interface {
	Add(context.Context, types.InputAddProduct) (int64, error)
	Delete(context.Context, int64) error
	GetByID(context.Context, int64) (types.OutputProduct, error)
	GetAll(context.Context, map[string]string, []string, int64, int64) ([]types.OutputProduct, error)
	UpdateByID(context.Context, int64, *types.InputUpdateProduct) error
}

// ProductService defines propertie
type ProductService struct {
	Storage storage.Storage
	Orm     helper.OrmInterface
}

// NewProductService map storage and return ProductService
func NewProductService() (ps ProductService) {
	ps.Storage = storage.NewStorage()
	ps.Orm = helper.NewOrm()
	return
}

// Add service for add a new product
func (ps ProductService) Add(ctx context.Context, product types.InputAddProduct) (id int64, err error) {
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
	ormer := ps.Orm.NewOrms()                            // Declare ormer
	ormer.BeginTx(ctx, &sql.TxOptions{})                 // Begin transaction
	id, err = ps.Storage.Product.Add(ormer, &inputModel) // Execute method Add
	if id < 1 || err != nil {
		ormer.Rollback()
		return
	}
	err = ormer.Commit()
	return
}

// Delete service for delete product by ID
func (ps ProductService) Delete(ctx context.Context, id int64) (err error) {
	modelProduct := models.Product{
		ID: id,
	}
	ormer := ps.Orm.NewOrms()
	num, err := ps.Storage.Product.Delete(ormer, &modelProduct)
	if num < 1 {
		errorMessage := "Not Found"
		err = errors.New(errorMessage)
	}
	return
}

// GetByID service retrieve product by ID
func (ps ProductService) GetByID(ctx context.Context, id int64) (result types.OutputProduct, err error) {
	ormer := ps.Orm.NewOrms()
	product, err := ps.Storage.Product.GetByID(ormer, id)
	if err != nil {
		return
	}
	copier.Copy(&result, &product)
	return
}

// GetAll service for retrieves all product matches certain condition
func (ps ProductService) GetAll(
	ctx context.Context,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (results []types.OutputProduct, err error) {
	ormer := ps.Orm.NewOrms()
	products, err := ps.Storage.Product.GetAll(ormer, query, order, offset, limit)
	copier.Copy(&results, &products)
	return
}

// UpdateByID service for update product by ID
func (ps ProductService) UpdateByID(ctx context.Context, id int64, product *types.InputUpdateProduct) (err error) {
	ormer := ps.Orm.NewOrms()
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
		return
	}
	return
}
