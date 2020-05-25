package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Product_20200515_161334 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Product_20200515_161334{}
	m.Created = "20200515_161334"

	migration.Register("Product_20200515_161334", m)
}

// Up Run the migrations
func (m *Product_20200515_161334) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE product(
		id serial primary key,
		name varchar(255) NOT NULL,
		detail TEXT NULL,
		category_id integer DEFAULT NULL,
		stock integer DEFAULT NULL,
		brand varchar(100) NULL,
		model varchar(255) NULL,
		price numeric NOT NULL,
		cost numeric NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
	)`)
}

// Down Reverse the migrationsvarchar
func (m *Product_20200515_161334) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE product")
	m.SQL("DELETE FROM migrations WHERE name = 'Product_20200515_161334'")
}
