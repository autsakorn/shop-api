package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Product represents all possible actions available to deal with data
type Product interface {
	Add(*models.Product) (int64, error)
	Delete(*models.Product) (int64, error)
	GetAll(map[string]string, []string, int64, int64) ([]models.Product, error)
	GetByID(int64) (models.Product, error)
	UpdateByID(*models.Product) (int64, error)
}

// ProductStorage defines properties
type ProductStorage struct{}

// NewProductStorage return ProductStorage
func NewProductStorage() (productStorage ProductStorage) { return }

// Add new record
func (s ProductStorage) Add(product *models.Product) (id int64, err error) {
	// Prepare data create
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(product)
	return
}

// Delete method delete product by ID
func (s ProductStorage) Delete(product *models.Product) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(product)
	return
}

// GetAll retrieves all Product matches certain condition. Returns empty list if
// no records exist
func (s ProductStorage) GetAll(
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (result []models.Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Product))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}
	qs = qs.OrderBy(order...)
	_, err = qs.All(&result)
	return
}

// GetByID method retrieves product by ID
func (s ProductStorage) GetByID(id int64) (v models.Product, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(models.Product)).Filter("id", id).RelatedSel().One(&v)
	return
}

// UpdateByID method update product by ID
func (s ProductStorage) UpdateByID(m *models.Product) (num int64, err error) {
	o := orm.NewOrm()
	m.UpdatedAt = time.Now()
	num, err = o.Update(m)
	return
}
