package main

import (
	"armada-api/cmd/server"
	"armada-api/databases"

	_ "armada-api/databases/migrations"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd [migrate|rollback|server]")
		return
	}

	command := os.Args[1]

	switch command {
		case "migrate":
			db := databases.InitDB()
			defer db.Close()
			databases.RunMigrations(db)

		case "migrate:rollback":
			db := databases.InitDB()
			defer db.Close()
			databases.RollbackMigrations(db)

		case "migration:create":
			if len(os.Args) < 3 {
				fmt.Println("Usage: go run index.go migration:create <migration_name>")
				return
			}
			name := os.Args[2]
			databases.CreateMigration(name)

		case "server":
			server.Run()

		default:
			fmt.Println("Unknown command:", command)
	}
}
