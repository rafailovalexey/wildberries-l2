package repository

import dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"

type RepositoryEventsInterface interface {
	CreateEvent(*dto.CreateEventDto) (*dto.EventDto, error)
	UpdateEvent(*dto.UpdateEventDto) (*dto.EventDto, error)

	GetEventsByUserIdAndPeriod(string, *dto.EventPeriodDto) (*dto.EventsDto, error)
}
