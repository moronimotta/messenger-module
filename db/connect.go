// db/connect.go
package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (Database, error) {
	dsn := buildDSNFromEnv()
	if dsn == "" {
		log.Fatalf("database configuration missing: set DB_URL or DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME in environment or .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(
		&UserModel{},
		&PlanModel{},
		&UserPlanModel{},
		&IntegrationModel{},
		&MessageModel{},
		&MessageStatusModel{},
	); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return &GormDatabase{DB: db}, nil
}

// buildDSNFromEnv returns a Postgres DSN string using environment variables.
// Priority: DB_URL if set; otherwise assemble from individual DB_* variables.
func buildDSNFromEnv() string {
	// Highest priority: full URL
	if url := os.Getenv("DB_URL"); url != "" {
		return url
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	// If any critical piece is missing, return empty to trigger a helpful fatal above
	if host == "" || port == "" || user == "" || name == "" {
		return ""
	}

	// Default sslmode to disable unless already specified via DB_SSLMODE or present in DB_URL
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	// GORM Postgres driver accepts both URL and key=value DSN styles
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, sslmode)
}
