package whatsapp

import (
	"context"
	"time"
)

func InitWhatsapp() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	InitStorage(ctx)
	NewWhatsappManager(&GlobalSessions)
	go func(ctx context.Context) {
		Manager.LoadDeviceFromStorage(ctx)
	}(ctx)
	defer cancel()
}
