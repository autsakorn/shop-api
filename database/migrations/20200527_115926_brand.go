package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Brand_20200527_115926 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Brand_20200527_115926{}
	m.Created = "20200527_115926"

	migration.Register("Brand_20200527_115926", m)
}

// Up Run the migrations
func (m *Brand_20200527_115926) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE brand(id serial primary key,title varchar(255) NOT NULL,slug TEXT NOT NULL,created_at TIMESTAMP WITH TIME ZONE NOT NULL,updated_at TIMESTAMP WITH TIME ZONE NOT NULL)")
}

// Down Reverse the migrations
func (m *Brand_20200527_115926) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE brand")
	m.SQL("DELETE FROM migrations WHERE name = 'Brand_20200527_115926'")
}
