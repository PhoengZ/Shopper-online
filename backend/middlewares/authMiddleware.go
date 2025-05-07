package middlewares

import (
	"backend/utils"
	"backend/utils/response"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
			return
		}
		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			response.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			return
		}
		token := tokenParts[1]
		userID, err := utils.ValidateToken(token)
		if err != nil {
			response.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			return
		}
		r.Header.Set("userID", userID)
		next.ServeHTTP(w, r)
	})
}
