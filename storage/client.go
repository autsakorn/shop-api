package storage

import (
	"shop-api/models"

	"github.com/astaxie/beego/orm"
)

// Client represents all possible actions available to deal with data
type Client interface {
	GetByXApiKeyAndStatus(orm.Ormer, string, int32) (int64, error)
}

// ClientStorage define properties ClientStorage
type ClientStorage struct{}

// NewClientStorage return ClientStorage
func NewClientStorage() (categoryStorage ClientStorage) {
	return
}

// GetByXApiKeyAndStatus method get by x_api_key and status
func (s ClientStorage) GetByXApiKeyAndStatus(ormer orm.Ormer, xAPIKey string, status int32) (id int64, err error) {
	result := models.Client{}
	err = ormer.QueryTable(new(models.Client)).
		Filter("x_api_key", xAPIKey).
		Filter("status", status).
		RelatedSel().One(&result)
	id = result.ID
	return
}
