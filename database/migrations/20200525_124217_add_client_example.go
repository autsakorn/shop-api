package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddClientExample_20200525_124217 struct {
	migration.Migration
}

type client = struct {
	Name    string `json:"name"`
	XApiKey string `json:"x_api_key"`
}

// DO NOT MODIFY
func init() {
	m := &AddClientExample_20200525_124217{}
	m.Created = "20200525_124217"

	migration.Register("AddClientExample_20200525_124217", m)
}

// Up Run the migrations
func (m *AddClientExample_20200525_124217) Up() {
	jsonFile, err := os.Open("../fixtures/client_initial.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]client, 0)
	json.Unmarshal(byteValue, &result)
	status := 1
	for _, value := range result {
		sql := fmt.Sprintf(`INSERT INTO "client" ("name", "x_api_key", "status") VALUES ('%s', '%s', '%d')`, value.Name, value.XApiKey, status)
		m.SQL(sql)
	}
}

// Down Reverse the migrations
func (m *AddClientExample_20200525_124217) Down() {
	jsonFile, err := os.Open("../fixtures/client_initial.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]client, 0)
	json.Unmarshal(byteValue, &result)

	for _, value := range result {
		sql := fmt.Sprintf(`DELETE FROM "client" WHERE "name" = '%s' AND "x_api_key" = '%s'`, value.Name, value.XApiKey)
		m.SQL(sql)
	}
	m.SQL("DELETE FROM migrations WHERE name = 'AddClientExample_20200525_124217'")
}
