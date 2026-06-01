package configdb

import (
	"database/sql"
	"log"
	"os"
)

func migrateDB(db *sql.DB) {
	query, err := os.ReadFile("./internal/db/migration/user_table.sql")
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}
}
