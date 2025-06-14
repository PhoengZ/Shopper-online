package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	r := mux.NewRouter()
	config.ConnectDB()
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterValidateRoutes(r)
	routes.RegisterCategoriesRoutes(r)
	routes.RegisterationTopUp(r)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(r)
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port ", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
