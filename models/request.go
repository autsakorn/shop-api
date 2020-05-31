package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Request defines properties
type Request struct {
	ID              int64     `orm:"column(id);auto"`
	Title           string    `orm:"column(title);size(255)"`
	Description     string    `orm:"column(description);size(128)"`
	Price           float64   `orm:"column(price)"`
	DueDate         time.Time `orm:"column(due_date);type(datetime)"`
	RequirePackage  string    `orm:"column(require_package);size(128)"`
	Remark          string    `orm:"column(remark);size(128)"`
	ShippingAddress string    `orm:"column(shipping_address);size(128)"`
	ShippingMethod  string    `orm:"column(shipping_method);size(255)"`
	Country         string    `orm:"column(country);size(255)"`
	CreatedAt       time.Time `orm:"column(created_at);type(datetime)"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(datetime)"`
	Category        *Category `orm:"column(category_id);rel(fk);null"`
	Brand           *Brand    `orm:"column(brand_id);rel(fk);null"`
}

func init() {
	orm.RegisterModel(new(Request))
}
