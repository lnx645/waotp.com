package database

import (
	"fmt"
	"log"
	"time"

	"dadandev.com/wa-engine/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Config config.Databse
	db     *sqlx.DB
}

var DB *Database

func InitDB(config config.Databse) {
	DB = &Database{
		Config: config,
	}
}

func (d *Database) Connect() error {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.Config.User, d.Config.Pass, d.Config.Host, d.Config.Port, d.Config.Name))

	if err != nil {
		return err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	d.db = db
	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}
	log.Println("Database terhubung: Pool koneksi berhasil dibuat.")
	return err
}

func (d *Database) GetConnection() *sqlx.DB {
	if d.db != nil {
		return d.db
	}
	return nil
}
