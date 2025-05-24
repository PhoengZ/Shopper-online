package routes

import (
	"backend/controllers"
	"net/http"
)

func RegisterValidateRoutes() {
	http.HandleFunc("/validate", controllers.ValidToken)
}
