package server

import (
	"encoding/json"
	"fmt"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/chain"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/interceptor"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/middleware"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"log"
	"net/http"
	"os"
)

func Run(controllerEvents controller.ControllerEventsInterface) {
	router := http.NewServeMux()

	middlewares := chain.ChainHandlers(
		interceptor.LoggingInterceptor,
		middleware.CorsMiddleware,
		middleware.AuthenticationMiddleware,
	)

	handler := http.HandlerFunc(controllerEvents.EventsHandler)
	wrapped := middlewares(handler)

	router.Handle("/v1/events/", wrapped)

	router.HandleFunc("/", NotFound)

	hostname := os.Getenv("HOSTNAME")

	port := os.Getenv("PORT")

	if port == "" {
		log.Panicf("specify the port")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("http server starts at address %s\n", address)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Panicf("error when starting the server %v\n", err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	WriteNotFoundError(w)

	return
}

func WriteNotFoundError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(SerializeError("not found"))
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
