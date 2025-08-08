package main

import (
	"log"
	"net/http"

	"github.com/LuckyGoyal039/airtable-repo/airtable"
	airtablegen "github.com/LuckyGoyal039/airtable-repo/api/airtable" // update this to actual import path
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	e := echo.New()
	airtableService := &airtable.AirtableService{}

	// Register handlers from generated OpenAPI code
	airtablegen.RegisterHandlers(e, airtableService)

	log.Println("Server running at http://localhost:8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// sample api: http://localhost:8080/airtable/appOsM0fcKAqWmxca/Contacts
