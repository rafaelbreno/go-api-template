package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DATABASE_URL = "postgresql://%s:%s@%s:%s/%s"
)

var DB *gorm.DB

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(postgres.Open(mountDatabaseURL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func mountDatabaseURL() string {
	return fmt.Sprintf(
		DATABASE_URL,
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}
