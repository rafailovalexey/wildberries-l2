package events

import (
	"encoding/json"
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
	switch request.Method {
	case http.MethodGet:
		switch request.RequestURI {
		case "/v1/events/events_for_day":
			c.EventsForDay(writer, request)
		case "/v1/events/events_for_week":
			c.EventsForWeek(writer, request)
		case "/v1/events/events_for_month":
			c.EventsForMonth(writer, request)
		default:
			WriteErrorNotFound(writer)
		}
	case http.MethodPost:
		switch request.RequestURI {
		case "/v1/events/create_event":
			c.CreateEvent(writer, request)
		case "/v1/events/update_event":
			c.UpdateEvent(writer, request)
		default:
			WriteErrorNotFound(writer)
		}
	default:
		WriteErrorMethodNotAllowed(writer)
	}
}

func WriteResultOk(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(SerializeResult(message))
}

func WriteResultCreated(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(SerializeResult(message))
}

func WriteErrorBadRequest(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write(SerializeError(message))
}

func WriteErrorMethodNotAllowed(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)
	writer.Write(SerializeError("method not allowed"))
}

func WriteErrorNotFound(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)
	writer.Write(SerializeError("not found"))
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
