package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DATABASE_URL = "postgresql://%s:%s@%s:%s/%s"
)

var DBConn *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open(mountDatabaseURL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DBConn = db
}

func mountDatabaseURL() string {
	return fmt.Sprintf(
		DATABASE_URL,
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}
