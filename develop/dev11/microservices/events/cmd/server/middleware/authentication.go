package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Panicf("specify the name of the authentication token")
		}

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Panicf("specify the value of the authentication token")
		}

		key := r.Header.Get(header)

		if key != token {
			WriteUnauthorizedError(w)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func WriteUnauthorizedError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(SerializeError("unauthorized"))
}

func SerializeError(message string) []byte {
	type Error struct {
		Error string `json:"error"`
	}

	e := &Error{
		Error: message,
	}

	j, err := json.Marshal(e)

	if err != nil {
		return []byte(err.Error())
	}

	return j
}
