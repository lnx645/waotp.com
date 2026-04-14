package api

import (
	"context"
	"fmt"
	"net/http"

	"dadandev.com/wa-engine/internal/network"
	"dadandev.com/wa-engine/pkg/whatsapp"
	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
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
			if ok := wa.GetClient().EnableAutoReconnect; ok {
				w.Write([]byte("OK"))

			}
			msg := &waE2E.Message{
				Conversation: proto.String("*123456* adalah kode verifikasi (OTP) Anda. Demi keamanan, JANGAN berikan kode ini kepada siapapun termasuk pihak [NamaAplikasi]. Kode ini hanya berlaku selama 5 menit."),
			}
			target := types.NewJID("628155292003", types.DefaultUserServer)
			SendLocation(wa.GetClient(), target)
			resp, _ := wa.GetClient().SendMessage(context.Background(), target, msg)
			fmt.Println(resp)
			w.Write([]byte(id))
		} else {
			w.Write([]byte("Disconnected"))
		}
	})
}

func SendLocation(client *whatsmeow.Client, jid types.JID) {
	// 1. Susun pesan lokasi
	msg := &waE2E.Message{
		LocationMessage: &waE2E.LocationMessage{
			DegreesLatitude:  proto.Float64(-6.9034212),
			DegreesLongitude: proto.Float64(107.4783446),
			Name:             proto.String("Bandung"),
			Address:          proto.String("Kota Bandung..."),
			IsLive:           proto.Bool(true),
			AccuracyInMeters: proto.Uint32(10),
		},
	}

	// 2. Kirim pesan
	resp, err := client.SendMessage(context.Background(), jid, msg)
	if err != nil {
		fmt.Printf("Gagal kirim lokasi: %v\n", err)
		return
	}

	fmt.Printf("Lokasi terkirim! ID: %s\n", resp.ID)
}
