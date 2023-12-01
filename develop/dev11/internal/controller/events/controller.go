package events

import (
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/controller"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/validation"
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
