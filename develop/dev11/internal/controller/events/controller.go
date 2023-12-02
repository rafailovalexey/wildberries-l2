package events

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"net/http"
	"time"
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
	createEventDto := &dto.CreateEventDto{}

	if err := json.NewDecoder(request.Body).Decode(createEventDto); err != nil {
		WriteErrorBadRequest(writer, err.Error())

		return
	}

	defer request.Body.Close()

	fmt.Println(createEventDto)
}

func (c *ControllerEvents) UpdateEvent(writer http.ResponseWriter, request *http.Request) {
	updateEventDto := &dto.UpdateEventDto{}

	if err := json.NewDecoder(request.Body).Decode(updateEventDto); err != nil {
		WriteErrorBadRequest(writer, err.Error())

		return
	}

	defer request.Body.Close()

	fmt.Println(updateEventDto)
}

func (c *ControllerEvents) EventsForDay(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForDayValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(writer, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForDayDto := dto.NewEventsForDayDto(
		userId,
		parsed,
	)

	fmt.Println(eventsForDayDto)
}

func (c *ControllerEvents) EventsForWeek(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForWeekValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(writer, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForWeekDto := dto.NewEventsForWeekDto(
		userId,
		parsed,
	)

	fmt.Println(eventsForWeekDto)
}

func (c *ControllerEvents) EventsForMonth(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForMonthValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(writer, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForMonthDto := dto.NewEventsForMonthDto(
		userId,
		parsed,
	)

	fmt.Println(eventsForMonthDto)
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
