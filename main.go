package main

import (
	"fmt"
	"github.com/HiBang15/golang-api-base-tulpo.git/api/rest/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	//run Rest API
	fmt.Println("Start run REST API .....")
	router.Start(os.Getenv("ENVIRONMENT"))
}