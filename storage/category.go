package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Category represents the action avilable
type Category interface {
	Add(input *models.Category) (id int64, err error)
	Delete(input *models.Category) (num int64, err error)
	GetAll(query map[string]string, fields []string, sortby []string, order []string,
		offset int64, limit int64) (result []models.Category, err error)
	GetByID(id int64) (result models.Category, err error)
	UpdateByID(input *models.Category) (num int64, err error)
}

// CategoryStorage ...
type CategoryStorage struct{}

// NewCategoryStorage ...
func NewCategoryStorage() (categoryStorage CategoryStorage) { return }

// Add ...
func (s CategoryStorage) Add(input *models.Category) (id int64, err error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(input)
	return
}

// Delete ...
func (s CategoryStorage) Delete(input *models.Category) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(input)
	return
}

// GetAll ...
func (s CategoryStorage) GetAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (result []models.Category, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Category))
	var ml []models.Category
	_, err = qs.All(&ml)
	for _, v := range ml {
		result = append(result, v)
	}
	return
}

// GetByID ...
func (s CategoryStorage) GetByID(id int64) (result models.Category, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(models.Category)).Filter("id", id).RelatedSel().One(&result)
	return
}

// UpdateByID ...
func (s CategoryStorage) UpdateByID(input *models.Category) (num int64, err error) {
	o := orm.NewOrm()
	input.UpdatedAt = time.Now()
	num, err = o.Update(input)
	return
}
