package events

import definition "github.com/emptyhopes/wildberries-l2-dev11/internal/validation"

type ValidationEvents struct {
}

var _ definition.ValidationEventsInterface = (*ValidationEvents)(nil)

func NewValidationEvents() *ValidationEvents {
	return &ValidationEvents{}
}
