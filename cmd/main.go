package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"streamflix-api/db"
	"streamflix-api/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}


	db.InitDB()


	router := routes.SetupRouter()


	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}
}
