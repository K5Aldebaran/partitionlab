package dsn

import (
	"fmt"
	"os"
)

func FromEnv() string {
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = os.Getenv("DB_PASSWORD")
	}
	dbname := os.Getenv("DB_NAME")

	// Default values if not set
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if pass == "" {
		pass = "123"
	}
	if dbname == "" {
		dbname = "dehydration"
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslMode)
}
