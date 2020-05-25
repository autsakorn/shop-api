package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Client_20200525_121200 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Client_20200525_121200{}
	m.Created = "20200525_121200"

	migration.Register("Client_20200525_121200", m)
}

// Up Run the migrations
func (m *Client_20200525_121200) Up() {
	m.SQL(`CREATE TABLE client(
		id serial primary key,
		name varchar(255) NOT NULL,
		x_api_key TEXT NOT NULL,
		status integer DEFAULT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
	)`)
}

// Down Reverse the migrations
func (m *Client_20200525_121200) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE client")
	m.SQL("DELETE FROM migrations WHERE name = 'Client_20200525_121200'")
}
