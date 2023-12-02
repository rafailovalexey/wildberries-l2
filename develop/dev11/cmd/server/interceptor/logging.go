package interceptor

import (
	"log"
	"net/http"
	"time"
)

func LoggingInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		next.ServeHTTP(writer, request)

		duration := time.Since(start)

		log.Printf("%s %s %s - %s %v\n", request.Method, request.URL.Path, request.RemoteAddr, request.UserAgent(), duration)
	})
}
