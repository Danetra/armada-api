package main

import (
	"armada-api/databases"
)

func main() {
	db := databases.InitDB()
	defer db.Close()

	// Jalankan migration
	databases.RunMigrations(db)
}
