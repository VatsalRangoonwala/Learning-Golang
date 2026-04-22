package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	products := []Product{{Name: "Shirt", Price: 249.80}, {Name: "Shoes", Price: 349.80}}

	jsonData, err := json.Marshal(products)
	if err!=nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(jsonData)
}
