package main

import (
	"encoding/json"
	"fmt"
)

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

type Product2 struct {
	Name      string `json:"name"`
	ProductID int64  `json:"product_id,string"`
}

func main() {

	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	p := &Product{}
	p2 := &Product2{}
	err := json.Unmarshal([]byte(data), p)
	err2 := json.Unmarshal([]byte(data), p2)
	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(p2)
	fmt.Println(*p)
	fmt.Println(p)
}
