package whatsapp

import (
	"fmt"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

type WhatsappEngine interface {
	Connect(name string)
	IsConnected() bool
}
type engine struct {
	Whatsapp  *whatsmeow.Client
	Container *sqlstore.Container
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
