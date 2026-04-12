package api

import (
	"net/http"

	"dadandev.com/wa-engine/internal/network"
	"dadandev.com/wa-engine/pkg/whatsapp"
	"github.com/gorilla/mux"
)

func ApiHandler(api *mux.Router) {
	api.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		wa := whatsapp.Manager.GetOrCreate("WKWKWKKW")
		if !wa.IsConnected() {
			wa.Connect("device_1")
		}
		network.SendToRoom("anonim", "foo", "Bar")
		w.Write([]byte("WKWKWKW"))
	})

}
