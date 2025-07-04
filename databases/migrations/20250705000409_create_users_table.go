package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	fmt.Println("users table created")
	return nil
}

func downCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE users`)
	if err != nil {
		return err
	}

	fmt.Println("users table dropped")
	return nil
}
