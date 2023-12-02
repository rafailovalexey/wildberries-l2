package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "*")
		writer.Header().Add("Access-Control-Allow-Credentials", "true")

		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(writer, request)
	})
}
