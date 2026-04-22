package main

import "encoding/json"

type Order struct {
	Product Product `json:"product"`
	User    User    `json:"user"`
}

func orderToJson(o Order) ([]byte, error) {
	return json.Marshal(o)
}
