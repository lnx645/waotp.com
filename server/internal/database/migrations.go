package database

import (
	"database/sql"
	"fmt"
	"log"

	"dadandev.com/wa-engine/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(d *config.Databse) {
	db, err := sql.Open("mysql", fmt.Sprintf(DB_FORMAT, d.User, d.Pass, d.Host, d.Port, d.Name))
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Gagal Koneksi Instance mIgration: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("Database is already up to date.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully!")
	}
	defer db.Close()

}
