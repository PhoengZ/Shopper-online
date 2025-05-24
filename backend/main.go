package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	config.ConnectDB()
	routes.RegisterUserRoutes()
	routes.RegisterProductRoutes()
	routes.RegisterValidateRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(http.DefaultServeMux)
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port ", port)
	http.ListenAndServe(":"+port, handler)
}
