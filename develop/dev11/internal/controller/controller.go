package controller

import "net/http"

type ControllerEventsInterface interface {
	CreateEvent(http.ResponseWriter, *http.Request)
	UpdateEvent(http.ResponseWriter, *http.Request)
	EventsForDay(http.ResponseWriter, *http.Request)
	EventsForWeek(http.ResponseWriter, *http.Request)
	EventsForMonth(http.ResponseWriter, *http.Request)
	EventsHandler(http.ResponseWriter, *http.Request)
}
