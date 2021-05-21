package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HiBang15/golang-api-base-tulpo.git/api/rest/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	//run Rest API
	fmt.Println("Start run REST API .....")
	router.Start(os.Getenv("ENVIRONMENT"))
}
