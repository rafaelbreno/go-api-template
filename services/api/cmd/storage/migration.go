package storage

func Migrator(entities ...interface{}) {
	for _, entity := range entities {
		DBConn.
			Statement.
			AutoMigrate(entity)
	}
}
