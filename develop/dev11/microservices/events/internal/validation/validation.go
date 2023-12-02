package validation

import (
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
)

type ValidationEventsInterface interface {
	CreateEventValidation(*request.CreateEventRequest) error
	UpdateEventValidation(*request.UpdateEventRequest) error

	EventsForDayValidation(string, string) error
	EventsForWeekValidation(string, string) error
	EventsForMonthValidation(string, string) error
}
