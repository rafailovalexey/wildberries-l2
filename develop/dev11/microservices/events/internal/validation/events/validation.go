package events

import (
	"fmt"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"regexp"
	"time"
)

type ValidationEvents struct{}

var _ definition.ValidationEventsInterface = (*ValidationEvents)(nil)

func NewValidationEvents() *ValidationEvents {
	return &ValidationEvents{}
}

func (v *ValidationEvents) CreateEventValidation(createEventRequest *request.CreateEventRequest) error {
	if err := isValidUuid(createEventRequest.UserId, "user_id"); err != nil {
		return err
	}

	if err := isValidDate(createEventRequest.Date, "0000-00-00", "date"); err != nil {
		return err
	}

	return nil
}

func (v *ValidationEvents) UpdateEventValidation(updateEventRequest *request.UpdateEventRequest) error {
	if err := isValidUuid(updateEventRequest.UserId, "user_id"); err != nil {
		return err
	}

	if err := isValidDate(updateEventRequest.Date, "0000-00-00", "date"); err != nil {
		return err
	}

	return nil
}

func (v *ValidationEvents) EventsForDayValidation(userId string, date string) error {
	if err := isValidUuid(userId, "user_id"); err != nil {
		return err
	}

	if err := isValidDate(date, "0000-00-00", "date"); err != nil {
		return err
	}

	return nil
}

func (v *ValidationEvents) EventsForWeekValidation(userId string, date string) error {
	if err := isValidUuid(userId, "user_id"); err != nil {
		return err
	}

	if err := isValidDate(date, "0000-00-00", "date"); err != nil {
		return err
	}

	return nil
}

func (v *ValidationEvents) EventsForMonthValidation(userId string, date string) error {
	if err := isValidUuid(userId, "user_id"); err != nil {
		return err
	}

	if err := isValidDate(date, "0000-00-00", "date"); err != nil {
		return err
	}

	return nil
}

func isValidUuid(uuid string, field string) error {
	result := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`).MatchString(uuid)

	if !result {
		return fmt.Errorf("%s is not uuid", field)
	}

	return nil
}

func isValidDate(date string, layout string, field string) error {
	_, err := time.Parse(layout, date)

	if err != nil {
		return fmt.Errorf("invalid %s field date format %s", field, err)
	}

	return nil
}
