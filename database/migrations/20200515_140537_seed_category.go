package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type SeedCategory_20200515_140537 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SeedCategory_20200515_140537{}
	m.Created = "20200515_140537"

	migration.Register("SeedCategory_20200515_140537", m)
}

type category = struct {
	Name string
}

// Up Run the migrations
func (m *SeedCategory_20200515_140537) Up() {
	jsonFile, err := os.Open("../fixtures/category.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]category, 0)
	json.Unmarshal(byteValue, &result)

	for _, value := range result {
		detail := fmt.Sprintf("%s detail", strings.Replace(value.Name, "'", "''", -1))
		sql := fmt.Sprintf("INSERT INTO category (name, detail) VALUES ('%s', '%s')", strings.Replace(value.Name, "'", "''", -1), detail)
		m.SQL(sql)
	}
}

// Down Reverse the migrations
func (m *SeedCategory_20200515_140537) Down() {
	jsonFile, err := os.Open("../fixtures/category.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]category, 0)
	json.Unmarshal(byteValue, &result)

	for _, value := range result {
		sql := fmt.Sprintf("DELETE FROM category WHERE name = '%s'", strings.Replace(value.Name, "'", "''", -1))
		m.SQL(sql)
	}
	m.SQL("DELETE FROM migrations WHERE name = 'SeedCategory_20200515_140537'")
}
