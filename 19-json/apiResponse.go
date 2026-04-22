package main

import "encoding/json"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func createResponse(success bool, message string, data interface{}) ([]byte, error) {
	res := Response{success, message, data}
	return json.Marshal(res)
}
