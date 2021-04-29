package storage

import (
	"context"
	"fmt"
)

func Migrator(queries ...string) {
	for _, val := range queries {
		_, err := DBConn.Exec(context.Background(), val)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
