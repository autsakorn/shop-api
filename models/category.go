package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Category defines strcuture
type Category struct {
	ID        int64     `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(128)"`
	Detail    string    `orm:"column(detail);type(longtext)"`
	Status    int32     `orm:"column(status)"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime)"`
}

var invalid string = "Invalid"

// CategoryStatus define category status
var CategoryStatus = []string{"Inactive", "Active"}

// StatusRes resolver data to output format
func (category *Category) StatusRes() string {
	if category.Status < 0 {
		return invalid
	}
	return CategoryStatus[category.Status]
}

func init() {
	orm.RegisterModel(new(Category))
}
