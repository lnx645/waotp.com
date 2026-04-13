package whatsapp

import (
	"context"
	"fmt"

	"dadandev.com/wa-engine/internal/config"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var GlobalContainer *sqlstore.Container

func InitStorage(ctx context.Context) error {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	store, err := sqlstore.New(ctx, "sqlite", fmt.Sprintf("file:data/%s.db?_pragma=foreign_keys(1)", config.Get().Whatsapp.StorageName), dbLog)
	if err != nil {
		return err
	}
	GlobalContainer = store
	return err
}
