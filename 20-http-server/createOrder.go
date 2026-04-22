package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Order struct {
	Products []Product `json:"products"`
	User     User      `json:"user"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func createResponse(success bool, message string, data interface{}) ([]byte, error) {
	res := Response{success, message, data}
	return json.Marshal(res)
}

func parseOrder(r *http.Request) (Order, error) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		return order, err
	}
	if len(order.Products) == 0 {
		return order, errors.New("no products provided")
	}
	return order, nil
}

func validateOrder(order Order) error {
	for _, p := range order.Products {
		if p.Name == "" || p.Price <= 0 {
			return errors.New("Invalid product data")
		}
	}
	if order.User.Name == "" || order.User.Email == "" {
		return errors.New("Invalid User data")
	}

	return nil
}

func sendErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	res, err := createResponse(false, message, nil)
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

func writeJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		sendErrorJSON(w, 405, "Method Not Allowed")
		return
	}

	order, err := parseOrder(r)
	if err != nil {
		sendErrorJSON(w, 400, err.Error())
		return
	}

	err = validateOrder(order)
	if err != nil {
		sendErrorJSON(w, 400, err.Error())
		return
	}

	res, err := createResponse(true, "Order Created", order)
	if err != nil {
		sendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	writeJSON(w, 200, res)
}
