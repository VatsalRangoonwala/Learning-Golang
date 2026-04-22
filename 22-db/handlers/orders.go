package handlers

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/services"
	"project/utils"

	"github.com/go-chi/chi/v5"
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

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}

	orders, err := services.GetOrders()
	if err != nil {
		utils.SendErrorJSON(w, 500, err.Error())
		return
	}

	res, err := utils.CreateResponse(true, "Orders fetched", orders)
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}

	utils.WriteJSON(w, 200, res)
}

func GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := services.GetOrder(id)
	if err != nil {
		utils.SendErrorJSON(w, 404, "Order not found")
		return
	}

	res, err := utils.CreateResponse(true, "Order fetched", order)
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}

	utils.WriteJSON(w, 200, res)
}
