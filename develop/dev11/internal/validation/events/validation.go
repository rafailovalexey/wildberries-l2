package events

import (
	"fmt"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

type ValidationEvents struct{}

var _ definition.ValidationEventsInterface = (*ValidationEvents)(nil)

func NewValidationEvents() *ValidationEvents {
	return &ValidationEvents{}
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

func isString(value string, field string, length int) error {
	if reflect.TypeOf(value).Kind() != reflect.String {
		return fmt.Errorf("%s is not string", field)
	}

	if utf8.RuneCountInString(value) > length {
		return fmt.Errorf("%s is string to big max length %d", field, length)
	}

	return nil
}

func isTime(t time.Time, field string) error {
	if reflect.TypeOf(t) != reflect.TypeOf(time.Time{}) {
		return fmt.Errorf("%s is not a time", field)
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
