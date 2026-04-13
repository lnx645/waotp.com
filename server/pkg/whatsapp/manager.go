package whatsapp

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"dadandev.com/wa-engine/internal/database"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
)

type WhatsappManager struct {
	mu        sync.RWMutex
	container *sqlstore.Container
	sessions  map[string]WhatsappEngine
}

func (w *WhatsappManager) GetOrCreate(name string) WhatsappEngine {
	w.mu.Lock()
	defer w.mu.Unlock()
	if sess, ok := w.sessions[name]; ok {
		return sess
	}
	engine := w.GetEngine()
	w.sessions[name] = engine
	return engine
}
func NewWhatsappManager() {
	Manager = &WhatsappManager{
		sessions: GlobalSessions,
	}
}
func (w *WhatsappManager) GetEngine() WhatsappEngine {
	return NewWhatsappEngine()
}

func (w *WhatsappManager) LoadDeviceFromStorage() {
	rows, err := database.DB.GetConnection().Query("SELECT id, phone, device_engine_id, device_uuid, status FROM devices")
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	defer rows.Close()
	chanel := make(chan struct{}, 1)
	for rows.Next() {
		wg.Add(1)
		var id int
		var phone, engineID, uuid, status string
		err := rows.Scan(&id, &phone, &engineID, &uuid, &status)
		if err != nil {
			log.Println("Error scanning:", err)
			continue
		}
		go func(uuid string) {
			defer wg.Done()
			chanel <- struct{}{}
			defer func() { <-chanel }()
			engine := w.GetOrCreate(uuid)
			if !engine.GetClient().IsConnected() {
				engine.NewClient(uuid)
				time.Sleep(2 * time.Second)
				engine.GetClient().SendPresence(context.Background(), types.PresenceAvailable)
			}
		}(uuid)
		minDelay := 1
		maxDelay := 3
		randomSeconds := rand.Intn(maxDelay-minDelay+1) + minDelay
		fmt.Printf("Device %s diload, nunggu %d detik buat berikutnya...\n", uuid, randomSeconds)
		time.Sleep(time.Duration(randomSeconds) * time.Second)
	}
	fmt.Println("Menunggu semua device selesai diproses...")
	wg.Wait()
	fmt.Println("Semua device yang valid telah berhasil diload.")
}
func (w *WhatsappManager) FullLogout(name string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	ctx := context.Background()
	sess, ok := w.sessions[name]
	if !ok {
		fmt.Printf("Session %s tidak ditemukan di memory\n", name)
		return
	}
	client := sess.GetClient()
	if client != nil {
		if client.IsConnected() {
			err := client.Logout(ctx)
			if err != nil {
				client.Disconnect()
			}
		}
	}
	delete(w.sessions, name)
}
