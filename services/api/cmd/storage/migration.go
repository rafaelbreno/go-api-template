package storage

func Migrator(entities ...interface{}) {
	DBConn.
		Statement.
		AutoMigrate(entities...)
}
