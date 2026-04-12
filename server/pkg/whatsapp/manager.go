package whatsapp

import "sync"

type WhatsappManager struct {
	mu       sync.RWMutex
	sessions map[string]WhatsappEngine
}

var Manager *WhatsappManager

func (w *WhatsappManager) GetOrCreate(name string) WhatsappEngine {
	w.mu.Lock()
	defer w.mu.Unlock()
	if sess, ok := w.sessions[name]; ok {
		return sess
	}
	engine := NewWhatsappEngine()
	w.sessions[name] = engine
	return engine
}
func NewWhatsappManager() {
	Manager = &WhatsappManager{
		sessions: make(map[string]WhatsappEngine, 400),
	}
}
