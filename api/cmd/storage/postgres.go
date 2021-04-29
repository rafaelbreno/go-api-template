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
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}
