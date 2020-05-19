package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddColumnCategory_20200518_041931 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnCategory_20200518_041931{}
	m.Created = "20200518_041931"

	migration.Register("AddColumnCategory_20200518_041931", m)
}

// Up Run the migrations
func (m *AddColumnCategory_20200518_041931) Up() {
	m.SQL(`ALTER TABLE category
		ADD COLUMN status integer NULL DEFAULT '1',
		ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW();`)
}

// Down Reverse the migrations
func (m *AddColumnCategory_20200518_041931) Down() {
	m.SQL(`ALTER TABLE category DROP COLUMN status, DROP COLUMN created_at, DROP COLUMN updated_at;`)
	m.SQL(`DELETE FROM migrations WHERE name = 'AddColumnCategory_20200518_041931'`)
}
