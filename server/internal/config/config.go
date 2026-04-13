package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Databse struct {
	Host string
	Pass string
	Name string
	Port string
	User string
}
type App struct {
	Port string
	Host string
}
type Whatsapp struct {
	StorageName string
}
type Config struct {
	Database Databse
	Whatsapp
	App
}

func Get() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		Whatsapp: Whatsapp{StorageName: os.Getenv("WA_STORAGE_NAME")},
		App: App{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: Databse{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
		},
	}
}
