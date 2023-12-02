package events

import (
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
	"time"
)

type ConverterEvents struct{}

var _ definition.ConverterEventsInterface = (*ConverterEvents)(nil)

func NewConverterEvents() *ConverterEvents {
	return &ConverterEvents{}
}

func (c *ConverterEvents) MapEventDtoToEventModel(eventDto *dto.EventDto) *model.EventModel {
	return &model.EventModel{
		UserId:    eventDto.UserId,
		Date:      eventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *ConverterEvents) MapEventsDtoToEventsModel(eventsDto *dto.EventsDto) *model.EventsModel {
	eventsModel := make(model.EventsModel, len(*eventsDto))

	for index, value := range *eventsDto {
		eventsModel[index] = *c.MapEventDtoToEventModel(&value)
	}

	return &eventsModel
}

func (c *ConverterEvents) MapCreateEventDtoToEventModel(createEventDto *dto.CreateEventDto) *model.EventModel {
	return &model.EventModel{
		UserId:    createEventDto.UserId,
		Date:      createEventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *ConverterEvents) MapUpdateEventDtoToEventModel(updateEventDto *dto.UpdateEventDto) *model.EventModel {
	return &model.EventModel{
		UserId:    updateEventDto.UserId,
		Date:      updateEventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *ConverterEvents) MapEventModelToEventDto(eventModel *model.EventModel) *dto.EventDto {
	return &dto.EventDto{
		UserId:    eventModel.UserId,
		Date:      eventModel.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *ConverterEvents) MapEventModelToCreateEventDto(eventModel *model.EventModel) *dto.CreateEventDto {
	return &dto.CreateEventDto{
		UserId: eventModel.UserId,
		Date:   eventModel.Date,
	}
}

func (c *ConverterEvents) MapEventModelToUpdateEventDto(eventModel *model.EventModel) *dto.UpdateEventDto {
	return &dto.UpdateEventDto{
		UserId: eventModel.UserId,
		Date:   eventModel.Date,
	}
}

func (c *ConverterEvents) MapEventsModelToEventsDto(eventsModel *model.EventsModel) *dto.EventsDto {
	eventsDto := make(dto.EventsDto, len(*eventsModel))

	for index, value := range *eventsModel {
		eventsDto[index] = *c.MapEventModelToEventDto(&value)
	}

	return &eventsDto
}
