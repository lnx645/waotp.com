package api

import (
	"dadandev.com/wa-engine/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterAuthRouter(api *mux.Router) {
	auth := handler.NewAuthHandler()
	api.HandleFunc("/login", auth.Login).Methods("POST")

}
