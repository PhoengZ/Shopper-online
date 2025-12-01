// package handler

// import (
// 	"backend/config"
// 	"backend/routes"
// 	"net/http"
// 	"os"
// 	"sync"

// 	"github.com/gorilla/mux"
// 	"github.com/rs/cors"
// )

// var router http.Handler
// var once sync.Once
// func init() {
// 	config.ConnectDB()
// 	r := mux.NewRouter()
// 	routes.RegisterUserRoutes(r)
// 	routes.RegisterProductRoutes(r)
// 	routes.RegisterValidateRoutes(r)
// 	routes.RegisterCategoriesRoutes(r)
// 	routes.RegisterationTopUp(r)

// 	// CORS
// 	domain := os.Getenv("FRONTEND_URL")
// 	if domain == "" {
// 		domain = "https://shopper-online-frontend.vercel.app"
// 	}
// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
// 		AllowedHeaders:   []string{"Content-Type", "Authorization"},
// 		AllowCredentials: false,
// 	})

// 	router = c.Handler(r)
// }

//	func Handler(w http.ResponseWriter, r *http.Request) {
//		router.ServeHTTP(w, r)
//	}
package handler

import (
	"backend/config"
	"backend/routes"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	router http.Handler
	once   sync.Once
)

func setup() {
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
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	})

	router = c.Handler(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setup)

	if router != nil {
		router.ServeHTTP(w, r)
	} else {
		http.Error(w, "Server Initialization Failed", http.StatusInternalServerError)
	}
}
