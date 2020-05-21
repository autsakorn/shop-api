package helper

import (
	"github.com/astaxie/beego/orm"
)

// NewOrm create and return new orm
func NewOrm(isTransaction bool) (ormer orm.Ormer) {
	ormer = orm.NewOrm()
	if isTransaction {
		ormer.Begin()
	}
	return
}
