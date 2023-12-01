package interceptor

import (
	"log"
	"net/http"
	"time"
)

func LoggingInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()

		next.ServeHTTP(response, request)

		duration := time.Since(start)

		log.Printf("%s %s %s - %s %v\n", request.Method, request.URL.Path, request.RemoteAddr, request.UserAgent(), duration)
	})
}
