package handlers

import (
	"net/http"
	"project/models"
	"project/utils"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}
	user := models.User{Name: "Vatsal", Email: "vatsal@gmail.com"}

	jsonData, err := utils.CreateResponse(true, "User Created", user)
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	utils.WriteJSON(w, 200, jsonData)
}
