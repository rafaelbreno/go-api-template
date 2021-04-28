package storage

import "context"

func Migrator(queries ...string) {
	for _, val := range queries {
		DBConn.Exec(context.Background(), val)
	}
}
