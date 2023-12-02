package server

import (
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

	hostname := os.Getenv("HOSTNAME")

	port := os.Getenv("PORT")

	if port == "" {
		log.Panicf("укажите порт")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("http сервер запускается по адресу %s\n", address)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Panicf("ошибка при запуске сервера: %v\n", err)
	}
}
