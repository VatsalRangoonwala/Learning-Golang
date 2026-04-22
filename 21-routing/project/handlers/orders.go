package handlers

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/services"
	"project/utils"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}

	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		utils.SendErrorJSON(w, 400, err.Error())
		return
	}

	processedOrder, err := services.ProcessOrder(order)
	if err != nil {
		utils.SendErrorJSON(w, 400, err.Error())
		return
	}

	res, err := utils.CreateResponse(true, "Order Created", processedOrder)
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	utils.WriteJSON(w, 200, res)
}
