package configdb

import (
	"banking-system/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func LoadDBConfig(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	migrateDB(sqlDB)
	return sqlDB
}
