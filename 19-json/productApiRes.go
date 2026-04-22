package main

import (
	"encoding/json"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func productApiResponse(products ...Product) ([]byte, error) {
	return json.Marshal(products)
}
