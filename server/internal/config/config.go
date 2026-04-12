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
type Config struct {
	Database Databse
	App
}

func Get() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		App: App{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
	}
}
