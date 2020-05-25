package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Client defines the properties of a client
type Client struct {
	ID        int64     `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255)"`
	XApiKey   string    `orm:"column(x_api_key);type(longtext)"`
	Status    int32     `orm:"column(status)"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Client))
}
