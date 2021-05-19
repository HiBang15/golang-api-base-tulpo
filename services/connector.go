package services

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var connDB *sql.DB

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Define connect to DB
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_USER"),
		os.Getenv("DB_USER_PASSWORD"), os.Getenv("DB_USER_HOST"),
		os.Getenv("DB_USER_PORT"), os.Getenv("DB_USER_NAME"), os.Getenv("DB_SSL_MODE"))
	connDB, err = sql.Open(os.Getenv("DB_DRIVER"), dbSource)
	if err != nil {
		log.Fatalf("has error occur when init connect to DB: %v", err)
	}
}
