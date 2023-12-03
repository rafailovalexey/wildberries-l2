package events

import (
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	response "github.com/emptyhopes/wildberries-l2-dev11/internal/response/events"
	"time"
)

type ConverterEvents struct{}

var _ definition.ConverterEventsInterface = (*ConverterEvents)(nil)

func NewConverterEvents() *ConverterEvents {
	return &ConverterEvents{}
}

func (c *ConverterEvents) MapCreateEventRequestToCreateEventDto(createEventRequest *request.CreateEventRequest) (*dto.CreateEventDto, error) {
	parsed, err := time.Parse("2006-01-02", createEventRequest.Date)

	if err != nil {
		return nil, err
	}

	createEventDto := dto.NewCreateEventDto(
		createEventRequest.UserId,
		parsed,
	)

	return createEventDto, nil
}

func (c *ConverterEvents) MapUpdateEventRequestToUpdateEventDto(updateEventRequest *request.UpdateEventRequest) (*dto.UpdateEventDto, error) {
	parsed, err := time.Parse("2006-01-02", updateEventRequest.Date)

	if err != nil {
		return nil, err
	}

	updateEventDto := dto.NewUpdateEventDto(
		updateEventRequest.Id,
		updateEventRequest.UserId,
		parsed,
	)

	return updateEventDto, nil
}

func (c *ConverterEvents) MapEventModelToEventDto(eventModel *model.EventModel) (*dto.EventDto, error) {
	eventDto := dto.NewEventDto(
		eventModel.Id,
		eventModel.UserId,
		eventModel.Date,
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
	)

	return eventDto, nil
}

func (c *ConverterEvents) MapEventsModelToEventsDto(eventsModel *model.EventsModel) (*dto.EventsDto, error) {
	eventsDto := make(dto.EventsDto, len(*eventsModel))

	for index, value := range *eventsModel {
		eventModel, err := c.MapEventModelToEventDto(&value)

		if err != nil {
			return nil, err
		}

		eventsDto[index] = *eventModel
	}

	return &eventsDto, nil
}

func (c *ConverterEvents) MapEventDtoToEventResponse(eventDto *dto.EventDto) (*response.EventResponse, error) {
	eventResponse := response.NewEventResponse(
		eventDto.Id,
		eventDto.UserId,
		eventDto.Date,
		eventDto.CreatedAt,
		eventDto.UpdatedAt,
	)

	return eventResponse, nil
}

func (c *ConverterEvents) MapEventsDtoToEventsResponse(eventsDto *dto.EventsDto) (*response.EventsResponse, error) {
	eventsResponse := make(response.EventsResponse, len(*eventsDto))

	for index, value := range *eventsDto {
		eventResponse, err := c.MapEventDtoToEventResponse(&value)

		if err != nil {
			return nil, err
		}

		eventsResponse[index] = *eventResponse
	}

	return &eventsResponse, nil
}
