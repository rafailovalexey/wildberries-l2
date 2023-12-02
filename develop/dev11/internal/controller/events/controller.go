package events

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"net/http"
)

type ControllerEvents struct {
	converterEvents  converter.ConverterEventsInterface
	validationEvents validation.ValidationEventsInterface
}

var _ definition.ControllerEventsInterface = (*ControllerEvents)(nil)

func NewControllerEvents(
	converterEvents converter.ConverterEventsInterface,
	validationEvents validation.ValidationEventsInterface,
) *ControllerEvents {
	return &ControllerEvents{
		converterEvents:  converterEvents,
		validationEvents: validationEvents,
	}
}

func (c *ControllerEvents) CreateEvent(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (c *ControllerEvents) UpdateEvent(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (c *ControllerEvents) EventsForDay(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (c *ControllerEvents) EventsForWeek(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (c *ControllerEvents) EventsForMonth(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (c *ControllerEvents) EventsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Println(request.RequestURI)
	fmt.Println(request.Method)

	switch {
	case request.Method == "POST" && request.RequestURI == "/v1/events/create_event":
		c.CreateEvent(writer, request)
	case request.Method == "POST" && request.RequestURI == "/v1/events/update_event":
		c.UpdateEvent(writer, request)
	case request.Method == "GET" && request.RequestURI == "/v1/events/events_for_day":
		c.EventsForDay(writer, request)
	case request.Method == "GET" && request.RequestURI == "/v1/events/events_for_week":
		c.EventsForWeek(writer, request)
	case request.Method == "GET" && request.RequestURI == "/v1/events/events_for_month":
		c.EventsForMonth(writer, request)
	default:
		http.Error(writer, "несуществующий метод", http.StatusMethodNotAllowed)
	}
}

func SerializeResult(message string) []byte {
	type Result struct {
		Result string `json:"result"`
	}

	e := &Result{
		Result: message,
	}

	j, err := json.Marshal(e)

	if err != nil {
		return []byte(err.Error())
	}

	return j
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
