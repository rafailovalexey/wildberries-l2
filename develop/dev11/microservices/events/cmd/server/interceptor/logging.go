package interceptor

import (
	"log"
	"net/http"
	"time"
)

func LoggingInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("%s %s %s - %s %v\n", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), duration)
	})
}
