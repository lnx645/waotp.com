package api

import (
	"net/http"

	"dadandev.com/wa-engine/internal/network"
	"dadandev.com/wa-engine/pkg/whatsapp"
	"github.com/gorilla/mux"
)

func ApiHandler(api *mux.Router) {
	api.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		wa := whatsapp.Manager.GetOrCreate(r.URL.Query().Get("name"))
		if !wa.IsConnected() {
			go wa.NewClient(r.URL.Query().Get("name"))
		}
		if wa.IsConnected() {
			network.SendToRoom("anonim", "foo", "Bar")
			id := wa.GetClient().Store.ID.String()
			w.Write([]byte(id))
		}
	})
	api.HandleFunc("/cek-koneksi", func(w http.ResponseWriter, r *http.Request) {
		wa := whatsapp.Manager.GetOrCreate(r.URL.Query().Get("name"))

		if wa.IsConnected() {
			id := wa.GetClient().Store.ID.String()
			w.Write([]byte(id))
		} else {
			w.Write([]byte("Disconnected"))
		}
	})
}
