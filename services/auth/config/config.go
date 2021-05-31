package config

import (
	"context"
	"fmt"
	"os"

	redis "github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	DATABASE_URL = "postgresql://%s:%s@%s:%s/%s"
)

var DB *gorm.DB
var Rdb *redis.Client

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

	if err != nil {
		panic(err)
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	if status := Rdb.Ping(context.Background()); status.Err() != nil {
		panic(status.Err())
	}

}

func mountDatabaseURL() string {
	return fmt.Sprintf(
		DATABASE_URL,
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}
