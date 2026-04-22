package main

import "net/http"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorJSON(w, 405, "Method Not Allowed")
		return
	}
	res, err := createResponse(true, "Healthy", map[string]string{"status": "ok"})
	if err != nil {
		sendErrorJSON(w, 500, "Internal Server Error")
		return
	}
	writeJSON(w, 200, res)
}
