package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Panicf("укажите имя токена аутентификации")
		}

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Panicf("укажите значение токена аутентификации")
		}

		key := request.Header.Get(header)

		if key != token {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusUnauthorized)
			response.Write(getErrorInJson("unauthorized"))

			return
		}

		next.ServeHTTP(response, request)
	})
}

func getErrorInJson(message string) []byte {
	type ErrorStruct struct {
		Error string `json:"error"`
	}

	errorStruct := &ErrorStruct{
		Error: message,
	}

	errJson, err := json.Marshal(errorStruct)

	if err != nil {
		return []byte(err.Error())
	}

	return errJson
}
