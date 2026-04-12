package whatsapp

import "fmt"

type WhatsappEngine interface {
	Connect(name string)
	IsConnected() bool
}
type engine struct {
}

func NewWhatsappEngine() *engine {
	return &engine{}
}
func (w *engine) Connect(name string) {
	fmt.Println(name)
}
func (w *engine) IsConnected() bool {
	return false
}
