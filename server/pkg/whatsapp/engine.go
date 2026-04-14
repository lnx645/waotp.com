package whatsapp

import (
	"context"
	"fmt"

	"dadandev.com/wa-engine/internal/config"
	"dadandev.com/wa-engine/internal/database"
	"dadandev.com/wa-engine/pkg/utils"
	_ "github.com/glebarez/go-sqlite"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type WhatsappEngine interface {
	NewClient(name string)
	IsConnected() bool

	Logout()
	GetJIDFromDB(name string) (string, error)
	GetClient() *whatsmeow.Client
}
type engine struct {
	Whatsapp  *whatsmeow.Client
	Container *sqlstore.Container
	storeName string
	ctx       context.Context
}

func (w *engine) Logout() {
	if w.Whatsapp.IsConnected() {
		err := w.Whatsapp.Logout(w.ctx)
		if err != nil {
			w.Whatsapp.Disconnect()
		}
	}
}
func NewWhatsappEngine() *engine {
	ctx := context.Background()
	conf := config.Get().Whatsapp
	return &engine{
		ctx:       ctx,
		storeName: conf.StorageName,
	}
}
func (w *engine) GetClient() *whatsmeow.Client {
	return w.Whatsapp
}
func (w *engine) updateDeviceStatus(client *whatsmeow.Client, sessionName string, status string) {
	// 1. Safety check: Pastikan store ID tidak nil (karena ini identitas utama)
	if client.Store.ID == nil {
		fmt.Println("Gagal update status: Store ID kosong (belum login/konek)")
		return
	}

	db := database.DB.GetConnection()
	phone := client.Store.ID.User
	deviceEngineID := client.Store.ID.String()
	query := `INSERT INTO devices (phone, device_engine_id, device_uuid, status, created_at, updated_at)
              VALUES (?, ?, ?, ?, NOW(), NOW())
              ON DUPLICATE KEY UPDATE
                device_engine_id = VALUES(device_engine_id),
                phone = VALUES(phone),
				status = VALUES(status),
                updated_at = NOW()`

	_, err := db.Exec(query, phone, deviceEngineID, sessionName, status)

	if err != nil {
		fmt.Printf("Gagal menyimpan status ke database untuk ID %s: %v\n", deviceEngineID, err)
	} else {
		fmt.Printf("Device %s [%s] berhasil di-%s!\n", phone, deviceEngineID, status)
	}
}
func (w *engine) GetJIDFromDB(uuid string) (string, error) {
	db := database.DB.GetConnection()
	var engineID string

	query := "SELECT device_engine_id FROM devices WHERE device_uuid = ? LIMIT 1"
	err := db.Get(&engineID, query, uuid)
	if err != nil {
		return "", err
	}
	return engineID, nil
}
func (w *engine) NewClient(name string) {
	w.storeName = name
	var device *store.Device
	uid, err := w.GetJIDFromDB(name)
	if err != nil {
		fmt.Println("Not device from db")
	}
	jid, err := types.ParseJID(uid)
	if err == nil {
		device, _ = GlobalContainer.GetDevice(w.ctx, jid)

	}
	if device == nil {
		fmt.Printf("[%s] Device tidak ditemukan, membuat session baru...\n", name)
		device = GlobalContainer.NewDevice()
	} else {
		fmt.Printf("[%s] Menggunakan session lama: %s\n", name, device.ID.String())
	}

	w.Container = GlobalContainer
	client := whatsmeow.NewClient(device, nil)
	client.AddEventHandler(func(evt any) {
		switch evt.(type) {
		case *events.Connected:
			go w.updateDeviceStatus(client, name, "connected")
			client.SendPresence(w.ctx, types.PresenceAvailable)
		case *events.Disconnected:
			go w.updateDeviceStatus(client, name, "disconnected")
		case *events.LoggedOut:
			go w.updateDeviceStatus(client, name, "logout")
		}
	})
	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			switch evt.Event {
			case "code":
				utils.GenerateAndSaveQrCode(evt.Code, fmt.Sprintf("public/qr/%s.png", name))
			case "timeout":
				fmt.Println("QR code:", "timeout")
			default:
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}
	w.Whatsapp = client
	fmt.Println(name)
}
func (w *engine) IsConnected() bool {
	return w.Whatsapp.IsConnected()
}

func (w *engine) IsLogout() bool {
	return !w.Whatsapp.IsLoggedIn()
}
