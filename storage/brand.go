package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Brand represents all possible actions available to deal with data
type Brand interface {
	Add(orm.Ormer, *models.Brand) (int64, error)
	Delete(orm.Ormer, *models.Brand) (int64, error)
	GetAll(orm.Ormer, map[string]string, []string, int64, int64) ([]models.Brand, error)
	GetByID(orm.Ormer, int64) (models.Brand, error)
	UpdateByID(orm.Ormer, *models.Brand) (int64, error)
}

// BrandStorage define properties BrandStorage
type BrandStorage struct{}

// NewBrandStorage return BrandStorage
func NewBrandStorage() (brandStorage BrandStorage) {
	return
}

// Add method add a new brand
func (s BrandStorage) Add(ormer orm.Ormer, input *models.Brand) (id int64, err error) {
	// Prepare data create
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	id, err = ormer.Insert(input)
	return
}

// Delete method delete a brand by ID
func (s BrandStorage) Delete(ormer orm.Ormer, input *models.Brand) (num int64, err error) {
	num, err = ormer.Delete(input)
	return
}

// GetAll retrieves all Brand matches certain condition. Returns empty list if
// no records exist
func (s BrandStorage) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (result []models.Brand, err error) {
	qs := ormer.QueryTable(new(models.Brand))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}
	qs = qs.OrderBy(order...)
	_, err = qs.RelatedSel().Limit(limit, offset).All(&result)
	return
}

// GetByID method retrieve all Brand match by ID
func (s BrandStorage) GetByID(ormer orm.Ormer, id int64) (result models.Brand, err error) {
	err = ormer.QueryTable(new(models.Brand)).Filter("id", id).RelatedSel().One(&result)
	return
}

// UpdateByID method update brand by ID
func (s BrandStorage) UpdateByID(ormer orm.Ormer, input *models.Brand) (num int64, err error) {
	input.UpdatedAt = time.Now()
	num, err = ormer.Update(input)
	return
}
