package server

import (
	"fmt"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/chain"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/interceptor"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server/middleware"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Run(controllerEvents controller.ControllerEventsInterface) {
	router := mux.NewRouter()

	middlewares := chain.ChainHandlers(
		interceptor.LoggingInterceptor,
		middleware.CorsMiddleware,
		middleware.AuthenticationMiddleware,
	)

	router.Use(middlewares)

	//router.NotFoundHandler = http.HandlerFunc(controllerEvents.NotFound)
	//router.MethodNotAllowedHandler = http.HandlerFunc(controllerEvents.MethodNotAllowed)

	//router.HandleFunc("/v1/employees/{id:[a-zA-Z0-9-]+}", controllerEvents.GetEmployeeById).Methods("GET")
	//router.HandleFunc("/v1/employees", controllerEvents.CreateEmployee).Methods("POST")

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
