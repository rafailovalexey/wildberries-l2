package events

import (
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/repository"
)

type RepositoryEvents struct {
	converterEvents converter.ConverterEventsInterface
}

var _ definition.RepositoryEventsInterface = (*RepositoryEvents)(nil)

func NewRepositoryEvents(
	converterEvents converter.ConverterEventsInterface,
) *RepositoryEvents {
	return &RepositoryEvents{
		converterEvents: converterEvents,
	}
}
