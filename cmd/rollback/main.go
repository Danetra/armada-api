package main

import (
	"armada-api/databases"
)

func main() {
	db := databases.InitDB()
	defer db.Close()

	// Jalankan rollback
	databases.RollbackMigrations(db)
}
