package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Panicf("specify the name of the authentication token")
		}

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Panicf("specify the value of the authentication token")
		}

		key := request.Header.Get(header)

		if key != token {
			WriteErrorUnauthorized(writer)

			return
		}

		next.ServeHTTP(writer, request)
	})
}

func WriteErrorUnauthorized(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	writer.Write(SerializeError("unauthorized"))
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
