package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Category_20200515_133557 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Category_20200515_133557{}
	m.Created = "20200515_133557"

	migration.Register("Category_20200515_133557", m)
}

// Run the migrations
func (m *Category_20200515_133557) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE category(id serial primary key,name TEXT NOT NULL, detail TEXT NOT NULL)")
}

// Reverse the migrations
func (m *Category_20200515_133557) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE category")
}
