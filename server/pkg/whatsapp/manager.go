package whatsapp

import "sync"

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
