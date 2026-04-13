package whatsapp

import (
	"context"
)

func InitWhatsapp() {
	ctx := context.Background()
	InitStorage(ctx)
	NewWhatsappManager()
	go Manager.LoadDeviceFromStorage()
}
