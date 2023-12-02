package events

import (
	"encoding/json"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/service"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"net/http"
	"strings"
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
		WriteBadRequestError(w, err.Error())

		return
	}

	defer r.Body.Close()

	err := c.validationEvents.CreateEventValidation(createEventRequest)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	createEventDto, err := c.converterEvents.MapCreateEventRequestToCreateEventDto(createEventRequest)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventDto, err := c.serviceEvents.CreateEvent(createEventDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventResponse, err := c.converterEvents.MapEventDtoToEventResponse(eventDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	WriteResultCreated(w, eventResponse)

	return
}

func (c *ControllerEvents) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	updateEventRequest := &request.UpdateEventRequest{}

	if err := json.NewDecoder(r.Body).Decode(updateEventRequest); err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	defer r.Body.Close()

	err := c.validationEvents.UpdateEventValidation(updateEventRequest)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	updateEventDto, err := c.converterEvents.MapUpdateEventRequestToUpdateEventDto(updateEventRequest)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventDto, err := c.serviceEvents.UpdateEvent(updateEventDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventResponse, err := c.converterEvents.MapEventDtoToEventResponse(eventDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	WriteResultOk(w, eventResponse)
}

func (c *ControllerEvents) EventsForDay(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForDayValidation(userId, date)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	parsed, _ := time.Parse("2006-01-02", date)

	eventsForDayDto := dto.NewEventsForDayDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForDay(eventsForDayDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	WriteResultOk(w, eventsResponse)
}

func (c *ControllerEvents) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForWeekValidation(userId, date)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	parsed, _ := time.Parse("2006-01-02", date)

	eventsForWeekDto := dto.NewEventsForWeekDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForWeek(eventsForWeekDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	WriteResultOk(w, eventsResponse)
}

func (c *ControllerEvents) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userId := query.Get("user_id")
	date := query.Get("date")

	err := c.validationEvents.EventsForMonthValidation(userId, date)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	parsed, _ := time.Parse("2006-01-02", date)

	eventsForMonthDto := dto.NewEventsForMonthDto(
		userId,
		parsed,
	)

	eventsDto, err := c.serviceEvents.GetEventsForMonth(eventsForMonthDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	eventsResponse, err := c.converterEvents.MapEventsDtoToEventsResponse(eventsDto)

	if err != nil {
		WriteBadRequestError(w, err.Error())

		return
	}

	WriteResultOk(w, eventsResponse)
}

func (c *ControllerEvents) EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch {
		case strings.Contains(r.RequestURI, "/v1/events/events_for_day"):
			c.EventsForDay(w, r)
		case strings.Contains(r.RequestURI, "/v1/events/events_for_week"):
			c.EventsForWeek(w, r)
		case strings.Contains(r.RequestURI, "/v1/events/events_for_month"):
			c.EventsForMonth(w, r)
		default:
			WriteNotFoundError(w)

			return
		}
	case http.MethodPost:
		switch {
		case strings.Contains(r.RequestURI, "/v1/events/create_event"):
			c.CreateEvent(w, r)
		case strings.Contains(r.RequestURI, "/v1/events/update_event"):
			c.UpdateEvent(w, r)
		default:
			WriteNotFoundError(w)

			return
		}
	default:
		WriteMethodNotAllowedError(w)

		return
	}
}

func WriteResultOk(w http.ResponseWriter, data interface{}) {
	result, err := SerializeResult(data)

	if err != nil {
		WriteInternalServerError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func WriteResultCreated(w http.ResponseWriter, data interface{}) {
	result, err := SerializeResult(data)

	if err != nil {
		WriteInternalServerError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func WriteBadRequestError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(SerializeError(message))
}

func WriteMethodNotAllowedError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(SerializeError("method not allowed"))
}

func WriteNotFoundError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(SerializeError("not found"))
}

func WriteInternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(SerializeError("internal server error"))
}

func SerializeResult(data interface{}) ([]byte, error) {
	type Result struct {
		Result interface{} `json:"result"`
	}

	e := &Result{
		Result: data,
	}

	j, err := json.Marshal(e)

	if err != nil {
		return nil, err
	}

	return j, nil
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
