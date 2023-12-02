package service

import (
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
)

type ServiceEventsInterface interface {
	CreateEvent(*dto.CreateEventDto) (*dto.EventDto, error)
	UpdateEvent(*dto.UpdateEventDto) (*dto.EventDto, error)
	GetEventsForDay(*dto.EventsForDayDto) (*dto.EventsDto, error)
	GetEventsForWeek(*dto.EventsForWeekDto) (*dto.EventsDto, error)
	GetEventsForMonth(*dto.EventsForMonthDto) (*dto.EventsDto, error)
}
