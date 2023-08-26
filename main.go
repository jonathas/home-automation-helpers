package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathas/home-automation-helpers/routes"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
		os.Exit(1)
  }

	router := routes.SetupRouter()
	router.Run()
}