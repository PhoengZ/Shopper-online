package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
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
	config.ConnectDB()
	routes.RegisterUserRoutes()
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port ", port)
	http.ListenAndServe(":"+port, nil)
}
