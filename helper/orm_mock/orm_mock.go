package ormmock

import (
	"context"
	"database/sql"
	"shop-api/helper"

	"github.com/astaxie/beego/orm"
)

// Ormer defines properties by used beego orm.Ormer
type Ormer struct {
	orm.Ormer
}

func newOrmer() Ormer {
	return Ormer{}
}

// BeginTx method for start transaction process
func (ormer Ormer) BeginTx(ctx context.Context, sqlTxOptions *sql.TxOptions) error {
	return nil
}

// Commit method for commit transaction process
func (ormer Ormer) Commit() error {
	return nil
}

// Rollback method for rollback transaction process
func (ormer Ormer) Rollback() error {
	return nil
}

// OrmMock defines properties empty
type OrmMock struct{}

// NewOrms method create new ormer
func (o OrmMock) NewOrms() helper.OrmerInterface {
	return newOrmer()
}

// RegisterDataBase mock beego RegisterDataBase return nil
func (o OrmMock) RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error {
	return nil
}
