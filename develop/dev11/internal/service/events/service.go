package events

import (
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
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

func (s *ServiceEvents) CreateEvent(createEventDto *dto.CreateEventDto) (*dto.EventDto, error) {
	eventDto, err := s.repositoryEvents.CreateEvent(createEventDto)

	if err != nil {
		return nil, err
	}

	return eventDto, nil
}

func (s *ServiceEvents) UpdateEvent(updateEventDto *dto.UpdateEventDto) (*dto.EventDto, error) {
	eventDto, err := s.repositoryEvents.UpdateEvent(updateEventDto)

	if err != nil {
		return nil, err
	}

	return eventDto, nil
}

func (s *ServiceEvents) GetEventsForDay(eventsForDayDto *dto.EventsForDayDto) (*dto.EventsDto, error) {
	eventPeriodDto := dto.NewEventPeriodDto(
		eventsForDayDto.Date,
		eventsForDayDto.Date,
	)

	eventsDto, err := s.repositoryEvents.GetEventsByUserIdAndPeriod(eventsForDayDto.UserId, eventPeriodDto)

	if err != nil {
		return nil, err
	}

	return eventsDto, nil
}

func (s *ServiceEvents) GetEventsForWeek(eventsForWeekDto *dto.EventsForWeekDto) (*dto.EventsDto, error) {
	eventPeriodDto := dto.NewEventPeriodDto(
		eventsForWeekDto.Date,
		eventsForWeekDto.Date,
	)

	eventsDto, err := s.repositoryEvents.GetEventsByUserIdAndPeriod(eventsForWeekDto.UserId, eventPeriodDto)

	if err != nil {
		return nil, err
	}

	return eventsDto, nil
}

func (s *ServiceEvents) GetEventsForMonth(eventsForMonthDto *dto.EventsForMonthDto) (*dto.EventsDto, error) {
	eventPeriodDto := dto.NewEventPeriodDto(
		eventsForMonthDto.Date,
		eventsForMonthDto.Date,
	)

	eventsDto, err := s.repositoryEvents.GetEventsByUserIdAndPeriod(eventsForMonthDto.UserId, eventPeriodDto)

	if err != nil {
		return nil, err
	}

	return eventsDto, nil
}
