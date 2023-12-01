package events

import (
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
)

type ConverterEvents struct{}

var _ definition.ConverterEventsInterface = (*ConverterEvents)(nil)

func NewConverterEvents() *ConverterEvents {
	return &ConverterEvents{}
}
