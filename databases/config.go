package databases

import (
	"armada-api/internal/model"
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Test ping
	if err := db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	DB = db
	return db
}

// RunMigrations executes goose up
func RunMigrations(db *sql.DB) {
	if err := goose.Up(db, "databases/migrations"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration applied successfully")
}

// RollbackMigrations rolls back the latest migration
func RollbackMigrations(db *sql.DB) {
	if err := goose.Down(db, "databases/migrations"); err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}
	log.Println("Rollback completed successfully")
}

func CreateMigration(name string) {
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("databases/migrations/%s_%s.go", timestamp, name)

	data := struct {
		FuncName      string
		TableName     string
		MigrationName string
	}{
		FuncName:      toCamelCase(name),
		TableName:     getTableName(name),
		MigrationName: name,
	}

	tmplContent, err := os.ReadFile("databases/templates/migration.go.tmpl")
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	tmpl, err := template.New("migration").Parse(string(tmplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	if err := os.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		fmt.Println("Error writing migration file:", err)
		return
	}

	fmt.Println("Migration created:", fileName)
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func getTableName(name string) string {
	if strings.HasPrefix(name, "create_") {
		parts := strings.Split(name[len("create_"):], "_")
		if len(parts) > 0 {
			return parts[0]
		}
	}
	return name
}

func InsertLocation(db *sql.DB, loc model.LocationPayload) {
	query := `INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, loc.VehicleID, loc.Latitude, loc.Longitude, loc.Timestamp)
	if err != nil {
		fmt.Println("DB insert error:", err)
	}
}