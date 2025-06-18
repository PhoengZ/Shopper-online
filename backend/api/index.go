package handler

import (
	"backend/config"
	"backend/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var router http.Handler

func init() {
	config.ConnectDB()
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterValidateRoutes(r)
	routes.RegisterCategoriesRoutes(r)
	routes.RegisterationTopUp(r)

	// CORS
	domain := os.Getenv("FRONTEND_URL")
	if domain == "" {
		domain = "https://shopper-online-frontend.vercel.app"
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{domain},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router = c.Handler(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
