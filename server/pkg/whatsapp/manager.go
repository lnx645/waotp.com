package whatsapp

import (
	"context"
	"fmt"
	"sync"
)

type WhatsappManager struct {
	mu       sync.RWMutex
	sessions map[string]WhatsappEngine
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
		sessions: make(map[string]WhatsappEngine, 400),
	}
}
func (w *WhatsappManager) GetEngine() WhatsappEngine {
	return NewWhatsappEngine()
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
