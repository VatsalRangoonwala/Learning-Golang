package handlers

import (
	"net/http"
	"project/models"
	"project/utils"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}
	products := []models.Product{{Name: "Shirt", Price: 249.80, Quantity: 2}, {Name: "Shoes", Price: 349.80, Quantity: 1}}

	jsonData, err := utils.CreateResponse(true, "Products Created", products)
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	utils.WriteJSON(w, 200, jsonData)
}
