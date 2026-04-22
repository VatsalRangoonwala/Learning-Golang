package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(success bool, message string, data interface{}) ([]byte, error) {
	res := Response{success, message, data}
	return json.Marshal(res)
}

func SendErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	res, err := CreateResponse(false, message, nil)
	if err != nil {
		fallback := []byte(`{"success":false,"message":"internal error"}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(fallback)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func WriteJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
