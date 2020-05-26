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
type ExampleProduct_20200525_072315 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ExampleProduct_20200525_072315{}
	m.Created = "20200525_072315"

	migration.Register("ExampleProduct_20200525_072315", m)
}

type product = struct {
	Name     string  `json:"name"`
	Detail   string  `json:"detail"`
	Category string  `json:"category"`
	Stock    int     `json:"stock"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Price    float64 `json:"price"`
	Cost     float64 `json:"cost"`
}

// Up Run the migrations
func (m *ExampleProduct_20200525_072315) Up() {
	jsonFile, err := os.Open("../fixtures/product_initial.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]product, 0)
	json.Unmarshal(byteValue, &result)

	for _, value := range result {
		categoryName := fmt.Sprintf("%s", strings.Replace(value.Category, "'", "''", -1))
		productName := fmt.Sprintf("%s", strings.Replace(value.Name, "'", "''", -1))
		productDetail := fmt.Sprintf("%s", strings.Replace(value.Detail, "'", "''", -1))
		sql := fmt.Sprintf(`INSERT INTO "product" (
				"name",
				"detail",
				"brand",
				"model",
				"stock",
				"price",
				"cost",
				"category_id"
			) VALUES (
				'%s',
				'%s',
				'%s',
				'%s',
				'%d',
				'%f',
				'%f',
				(SELECT "category"."id" FROM "category" WHERE "category"."name" = '%s')
			)`,
			productName,
			productDetail,
			value.Brand,
			value.Model,
			value.Stock,
			value.Price,
			value.Cost,
			categoryName,
		)
		m.SQL(sql)
	}
}

// Down Reverse the migrations
func (m *ExampleProduct_20200525_072315) Down() {
	jsonFile, err := os.Open("../fixtures/product_initial.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := make([]product, 0)
	json.Unmarshal(byteValue, &result)

	for _, value := range result {
		productName := strings.Replace(value.Name, "'", "''", -1)
		sql := fmt.Sprintf("DELETE FROM product WHERE name = '%s'", productName)
		m.SQL(sql)
	}
	m.SQL("DELETE FROM migrations WHERE name = 'ExampleProduct_20200525_072315'")

}
