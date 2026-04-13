package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"dadandev.com/wa-engine/internal/api"
	"dadandev.com/wa-engine/internal/config"
	"dadandev.com/wa-engine/internal/database"
	"dadandev.com/wa-engine/internal/middleware"
	"dadandev.com/wa-engine/internal/network"
	"dadandev.com/wa-engine/pkg/whatsapp"
)

func main() {
	end := make(chan os.Signal, 1)
	conf := config.Get()
	//init database
	database.InitDB(conf.Database)
	err := database.DB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	whatsapp.InitWhatsapp()
	r := network.InitServer()
	r.Use(middleware.Logger)
	a := r.PathPrefix("/api").Subrouter()
	api.ApiHandler(a)
	func() {
		fmt.Println("Running")
		if err := http.ListenAndServe(conf.Port, r); err != nil {
			log.Fatal(err.Error())
		}
	}()

	signal.Notify(end, os.Interrupt)
	<-end
	log.Println("shutting down")
}
