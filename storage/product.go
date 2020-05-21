package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Product represents all possible actions available to deal with data
type Product interface {
	Add(orm.Ormer, *models.Product) (int64, error)
	Delete(orm.Ormer, *models.Product) (int64, error)
	GetAll(orm.Ormer, map[string]string, []string, int64, int64) ([]models.Product, error)
	GetByID(orm.Ormer, int64) (models.Product, error)
	UpdateByID(orm.Ormer, *models.Product) (int64, error)
}

// ProductStorage defines properties
type ProductStorage struct{}

// NewProductStorage return ProductStorage
func NewProductStorage() (productStorage ProductStorage) { return }

// Add new record
func (s ProductStorage) Add(ormer orm.Ormer, product *models.Product) (id int64, err error) {
	// Prepare data create
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	id, err = ormer.Insert(product)
	return
}

// Delete method delete product by ID
func (s ProductStorage) Delete(ormer orm.Ormer, product *models.Product) (num int64, err error) {
	num, err = ormer.Delete(product)
	return
}

// GetAll retrieves all Product matches certain condition. Returns empty list if
// no records exist
func (s ProductStorage) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (result []models.Product, err error) {
	qs := ormer.QueryTable(new(models.Product))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}
	qs = qs.OrderBy(order...)
	_, err = qs.All(&result)
	return
}

// GetByID method retrieves product by ID
func (s ProductStorage) GetByID(ormer orm.Ormer, id int64) (v models.Product, err error) {
	err = ormer.QueryTable(new(models.Product)).Filter("id", id).RelatedSel().One(&v)
	return
}

// UpdateByID method update product by ID
func (s ProductStorage) UpdateByID(ormer orm.Ormer, m *models.Product) (num int64, err error) {
	m.UpdatedAt = time.Now()
	num, err = ormer.Update(m)
	return
}
