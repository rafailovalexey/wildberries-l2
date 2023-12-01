package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Access-Control-Allow-Origin", "*")
		response.Header().Add("Access-Control-Allow-Headers", "*")
		response.Header().Add("Access-Control-Allow-Methods", "*")
		response.Header().Add("Access-Control-Allow-Credentials", "true")

		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(response, request)
	})
}
