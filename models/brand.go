package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Brand defines properties
type Brand struct {
	ID        int64     `orm:"column(id);auto"`
	Title     string    `orm:"size(255)"`
	Slug      string    `orm:"type(longtext)"`
	CreatedAt time.Time `orm:"type(datetime)"`
	UpdatedAt time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Brand))
}
