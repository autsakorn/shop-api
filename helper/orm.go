package helper

import (
	"github.com/astaxie/beego/orm"
)

// OrmInterface defines possible actions available
type OrmInterface interface {
	NewOrms() OrmerInterface
	RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error
}

// Orm defines property
type Orm struct{}

// NewOrm func for return object orm
func NewOrm() Orm {
	o := Orm{}
	return o
}

// NewOrms method create ormer and return
func (o Orm) NewOrms() OrmerInterface {
	ormer := newOrmer()
	return ormer
}

// RegisterDataBase method for RegisterDataBase
func (o Orm) RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error {
	err := orm.RegisterDataBase(aliasName, driverName, dataSource, params...)
	return err
}

// Start defines Ormer

// OrmerInterface represents all possible actions available for Orm
type OrmerInterface interface {
	orm.Ormer
}

// Ormer defines properties ormer
type Ormer struct {
	orm.Ormer
}

func newOrmer() Ormer {
	ormer := orm.NewOrm()
	return Ormer{ormer}
}
