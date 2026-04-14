package utils

import (
	"encoding/json"
	"net/http"
)

func send(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)

}
func ErrorResponse(w http.ResponseWriter, status int, msg string) {
	send(w, http.StatusInternalServerError, struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: msg,
		Success: false,
	})
}
