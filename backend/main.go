package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	routes.RegisterUserRoutes()
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port ", port)
	http.ListenAndServe(":"+port, nil)
}
