package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateVehiclesTable, downCreateVehiclesTable)
}

func upCreateVehiclesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE vehicles (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			vehicle_id TEXT NOT NULL,
			latitude TEXT NOT NULL,
			longitude TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	fmt.Println("vehicles table created")
	return nil
}

func downCreateVehiclesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE vehicles`)
	if err != nil {
		return err
	}

	fmt.Println("vehicles table dropped")
	return nil
}
