package handlers

import (
	"net/http"
	"project/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello API")

	if r.Method != http.MethodGet {
		utils.SendErrorJSON(w, 405, "Method Not Allowed")
		return
	}
	utils.WriteJSON(w, 200, []byte("Hello from API"))
}
