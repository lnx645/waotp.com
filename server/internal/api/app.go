package api

import (
	"net/http"

	"dadandev.com/wa-engine/internal/network"
	"dadandev.com/wa-engine/pkg/whatsapp"
	"github.com/gorilla/mux"
)

func ApiHandler(api *mux.Router) {
	api.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		wa := whatsapp.Manager.GetOrCreate("device_1")
		if !wa.IsConnected() {
			go wa.Connect("device_1")
		}
		if wa.IsConnected() {
			network.SendToRoom("anonim", "foo", "Bar")
			id := wa.GetClient().Store.ID.String()
			w.Write([]byte(id))
		}
	})

}
