package storage

import (
	"shop-api/models"
	"time"

	"github.com/astaxie/beego/orm"
)

// Request represents all possible actions available to deal with data
type Request interface {
	Add(orm.Ormer, *models.Request) (int64, error)
	GetAll(orm.Ormer, map[string]string, []string, int64, int64) ([]models.Request, error)
	GetByID(orm.Ormer, int64) (models.Request, error)
	UpdateByID(orm.Ormer, *models.Request) (int64, error)
}

// RequestStorage define properties RequestStorage
type RequestStorage struct{}

// NewRequestStorage return RequestStorage
func NewRequestStorage() (requestStorage RequestStorage) {
	return
}

// Add method add a new request
func (s RequestStorage) Add(ormer orm.Ormer, input *models.Request) (id int64, err error) {
	// Prepare data create
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	id, err = ormer.Insert(input)
	return
}

// GetAll retrieves all Request matches certain condition. Returns empty list if
// no records exist
func (s RequestStorage) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	request []string,
	offset int64,
	limit int64,
) (result []models.Request, err error) {
	qs := ormer.QueryTable(new(models.Request))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}
	qs = qs.OrderBy(request...)
	_, err = qs.RelatedSel().Limit(limit, offset).All(&result)
	return
}

// GetByID method retrieve all Request match by ID
func (s RequestStorage) GetByID(ormer orm.Ormer, id int64) (result models.Request, err error) {
	err = ormer.QueryTable(new(models.Request)).Filter("id", id).RelatedSel().One(&result)
	return
}

// UpdateByID method update request by ID
func (s RequestStorage) UpdateByID(ormer orm.Ormer, input *models.Request) (num int64, err error) {
	input.UpdatedAt = time.Now()
	num, err = ormer.Update(input)
	return
}
