package events

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/service"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"net/http"
	"time"
)

type ControllerEvents struct {
	validationEvents validation.ValidationEventsInterface
	serviceEvents    service.ServiceEventsInterface
	converterEvents  converter.ConverterEventsInterface
}

var _ definition.ControllerEventsInterface = (*ControllerEvents)(nil)

func NewControllerEvents(
	validationEvents validation.ValidationEventsInterface,
	serviceEvents service.ServiceEventsInterface,
	converterEvents converter.ConverterEventsInterface,
) *ControllerEvents {
	return &ControllerEvents{
		validationEvents: validationEvents,
		serviceEvents:    serviceEvents,
		converterEvents:  converterEvents,
	}
}

func (c *ControllerEvents) CreateEvent(w http.ResponseWriter, r *http.Request) {
	createEventRequest := &request.CreateEventRequest{}

	if err := json.NewDecoder(r.Body).Decode(createEventRequest); err != nil {
		WriteErrorBadRequest(w, err.Error())

		return
	}

	defer r.Body.Close()

	err := c.validationEvents.CreateEventValidation(createEventRequest)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	createEventDto, err := c.converterEvents.MapCreateEventRequestToCreateEventDto(createEventRequest)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventDto, err := c.serviceEvents.CreateEvent(createEventDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventResponse, err := c.converterEvents.MapEventDtoToEventResponse(eventDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	fmt.Println(eventResponse)

	WriteResultCreated(w, "")
}

func (c *ControllerEvents) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	updateEventRequest := &request.UpdateEventRequest{}

	if err := json.NewDecoder(r.Body).Decode(updateEventRequest); err != nil {
		WriteErrorBadRequest(w, err.Error())

		return
	}

	defer r.Body.Close()

	err := c.validationEvents.UpdateEventValidation(updateEventRequest)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	updateEventDto, err := c.converterEvents.MapUpdateEventRequestToUpdateEventDto(updateEventRequest)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventDto, err := c.serviceEvents.UpdateEvent(updateEventDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventResponse, err := c.converterEvents.MapEventDtoToEventResponse(eventDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	fmt.Println(eventResponse)

	WriteResultOk(w, "")
}

func (c *ControllerEvents) EventsForDay(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForDayValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForDayDto := dto.NewEventsForDayDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForDay(eventsForDayDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	fmt.Println(eventsResponse)

	WriteResultOk(w, "")
}

func (c *ControllerEvents) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForWeekValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForWeekDto := dto.NewEventsForWeekDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForWeek(eventsForWeekDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	fmt.Println(eventsResponse)

	WriteResultOk(w, "")
}

func (c *ControllerEvents) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForMonthValidation(userId, date)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	parsed, _ := time.Parse("0000-00-00", date)

	eventsForMonthDto := dto.NewEventsForMonthDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForMonth(eventsForMonthDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteErrorBadRequest(w, err.Error())
	}

	fmt.Println(eventsResponse)

	WriteResultOk(w, "")
}

func (c *ControllerEvents) EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.RequestURI {
		case "/v1/events/events_for_day":
			c.EventsForDay(w, r)
		case "/v1/events/events_for_week":
			c.EventsForWeek(w, r)
		case "/v1/events/events_for_month":
			c.EventsForMonth(w, r)
		default:
			WriteErrorNotFound(w)
		}
	case http.MethodPost:
		switch r.RequestURI {
		case "/v1/events/create_event":
			c.CreateEvent(w, r)
		case "/v1/events/update_event":
			c.UpdateEvent(w, r)
		default:
			WriteErrorNotFound(w)
		}
	default:
		WriteErrorMethodNotAllowed(w)
	}
}

func WriteResultOk(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(SerializeResult(message))
}

func WriteResultCreated(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(SerializeResult(message))
}

func WriteErrorBadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(SerializeError(message))
}

func WriteErrorMethodNotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(SerializeError("method not allowed"))
}

func WriteErrorNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(SerializeError("not found"))
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
