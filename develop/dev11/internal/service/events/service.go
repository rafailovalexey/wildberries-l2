package events

import (
	"github.com/emptyhopes/wildberries-l2-dev11/internal/repository"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/service"
)

type ServiceEvents struct {
	repositoryEvents repository.RepositoryEventsInterface
}

var _ definition.ServiceEventsInterface = (*ServiceEvents)(nil)

func NewServiceEvents(
	repositoryEvents repository.RepositoryEventsInterface,
) *ServiceEvents {
	return &ServiceEvents{
		repositoryEvents: repositoryEvents,
	}
}
