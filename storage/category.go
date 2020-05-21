package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Category represents all possible actions available to deal with data
type Category interface {
	Add(ormer orm.Ormer, input *models.Category) (id int64, err error)
	Delete(ormer orm.Ormer, input *models.Category) (num int64, err error)
	GetAll(ormer orm.Ormer, query map[string]string, order []string, offset int64, limit int64) (result []models.Category, err error)
	GetByID(ormer orm.Ormer, id int64) (result models.Category, err error)
	UpdateByID(ormer orm.Ormer, input *models.Category) (num int64, err error)
}

// CategoryStorage define properties CategoryStorage
type CategoryStorage struct{}

// NewCategoryStorage return CategoryStorage
func NewCategoryStorage() (categoryStorage CategoryStorage) {
	return
}

// Add method add a new category
func (s CategoryStorage) Add(ormer orm.Ormer, input *models.Category) (id int64, err error) {
	// Prepare data create
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	id, err = ormer.Insert(input)
	return
}

// Delete method delete a category by ID
func (s CategoryStorage) Delete(ormer orm.Ormer, input *models.Category) (num int64, err error) {
	num, err = ormer.Delete(input)
	return
}

// GetAll retrieves all Category matches certain condition. Returns empty list if
// no records exist
func (s CategoryStorage) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (result []models.Category, err error) {
	qs := ormer.QueryTable(new(models.Category))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}
	qs = qs.OrderBy(order...)
	_, err = qs.All(&result)
	return
}

// GetByID method retrieve all Category match by ID
func (s CategoryStorage) GetByID(ormer orm.Ormer, id int64) (result models.Category, err error) {
	err = ormer.QueryTable(new(models.Category)).Filter("id", id).RelatedSel().One(&result)
	return
}

// UpdateByID method update category by ID
func (s CategoryStorage) UpdateByID(ormer orm.Ormer, input *models.Category) (num int64, err error) {
	input.UpdatedAt = time.Now()
	num, err = ormer.Update(input)
	return
}
