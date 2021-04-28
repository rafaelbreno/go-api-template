package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	DATABASE_URL = "postgresql://%s:%s@%s:%s/%s"
)

var DBConn *pgx.Conn

func init() {
	conn, err := pgx.Connect(context.Background(), mountDatabaseURL())
	if err != nil {
		panic(err)
	}

	DBConn = conn
}

func mountDatabaseURL() string {
	return fmt.Sprintf(
		DATABASE_URL,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}
