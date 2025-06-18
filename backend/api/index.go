package api

import (
	"backend/config"
	"backend/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB()

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterValidateRoutes(router)
	routes.RegisterCategoriesRoutes(router)
	routes.RegisterationTopUp(router)

	domain := os.Getenv("FRONTEND_URL")
	if domain == "" {
		domain = "https://shopper-online-frontend.vercel.app"
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{domain},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	c.Handler(router).ServeHTTP(w, r)
}
