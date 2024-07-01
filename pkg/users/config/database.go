package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type dbconfig struct {
	Driver string
	Source string
}

func DataBase() *dbconfig {
	_, filename, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(filename)
	envPath := filepath.Join(configDir, "../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("DataBaseConfig, error reading .env file:", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
	return &dbconfig{
		Driver: "postgres",
		Source: url,
	}
}
