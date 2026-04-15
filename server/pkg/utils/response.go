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
	send(w, status, struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: msg,
		Success: false,
	})
}
func ResponseErrorWithData(w http.ResponseWriter, status int, data any) {

	send(w, status, map[string]any{"errors": data, "status": status, "success": false})
}
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	send(w, http.StatusOK, data)
}
