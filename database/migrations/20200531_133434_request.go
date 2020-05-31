package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Request_20200531_133434 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Request_20200531_133434{}
	m.Created = "20200531_133434"

	migration.Register("Request_20200531_133434", m)
}

// Up Run the migrations
func (m *Request_20200531_133434) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE request(id serial primary key,title varchar(255) NOT NULL,description TEXT NOT NULL,price numeric NOT NULL,category_id integer DEFAULT NULL,brand_id integer DEFAULT NULL,due_date TIMESTAMP WITH TIME ZONE NOT NULL,require_package TEXT NOT NULL,remark TEXT NOT NULL,shipping_address TEXT NOT NULL,shipping_method varchar(255) NOT NULL,country varchar(255) NOT NULL,created_at TIMESTAMP WITH TIME ZONE NOT NULL,updated_at TIMESTAMP WITH TIME ZONE NOT NULL)")
}

// Down Reverse the migrations
func (m *Request_20200531_133434) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE request")
	m.SQL("DELETE FROM migrations WHERE name = 'Request_20200531_133434'")
}
