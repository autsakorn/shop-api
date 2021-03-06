package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Product defines the properties of a product
type Product struct {
	ID        int64     `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255)"`
	Detail    string    `orm:"column(detail);type(longtext)"`
	Brand     string    `orm:"column(brand);size(100)"`
	Model     string    `orm:"column(model);size(255)"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime)"`
	Stock     int       `orm:"column(stock)"`
	Price     float64   `orm:"column(price)"`
	Cost      float64   `orm:"column(cost)"`
	Category  *Category `orm:"column(category_id);rel(fk);null"`
}

func init() {
	orm.RegisterModel(new(Product))
}
