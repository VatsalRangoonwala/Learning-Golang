package handlers

import (
	"net/http"
	"project/utils"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}
	res, err := utils.CreateResponse(true, "Healthy", map[string]string{"status": "ok"})
	if err != nil {
		utils.SendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	utils.WriteJSON(w, 200, res)
}
