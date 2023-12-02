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

	createEventDto := &dto.CreateEventDto{
		UserId: createEventRequest.UserId,
		Date:   parsed,
	}

	return createEventDto, nil
}

func (c *ConverterEvents) MapUpdateEventRequestToUpdateEventDto(updateEventRequest *request.UpdateEventRequest) (*dto.UpdateEventDto, error) {
	parsed, err := time.Parse("2006-01-02", updateEventRequest.Date)

	if err != nil {
		return nil, err
	}

	updateEventDto := &dto.UpdateEventDto{
		UserId: updateEventRequest.UserId,
		Date:   parsed,
	}

	return updateEventDto, nil
}

func (c *ConverterEvents) MapEventDtoToEventModel(eventDto *dto.EventDto) (*model.EventModel, error) {
	eventModel := &model.EventModel{
		UserId:    eventDto.UserId,
		Date:      eventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return eventModel, nil
}

func (c *ConverterEvents) MapEventsDtoToEventsModel(eventsDto *dto.EventsDto) (*model.EventsModel, error) {
	eventsModel := make(model.EventsModel, len(*eventsDto))

	for index, value := range *eventsDto {
		eventModel, err := c.MapEventDtoToEventModel(&value)

		if err != nil {
			return nil, err
		}

		eventsModel[index] = *eventModel
	}

	return &eventsModel, nil
}

func (c *ConverterEvents) MapCreateEventDtoToEventModel(createEventDto *dto.CreateEventDto) (*model.EventModel, error) {
	eventModel := &model.EventModel{
		UserId:    createEventDto.UserId,
		Date:      createEventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return eventModel, nil
}

func (c *ConverterEvents) MapUpdateEventDtoToEventModel(updateEventDto *dto.UpdateEventDto) (*model.EventModel, error) {
	eventModel := &model.EventModel{
		UserId:    updateEventDto.UserId,
		Date:      updateEventDto.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return eventModel, nil
}

func (c *ConverterEvents) MapEventModelToEventDto(eventModel *model.EventModel) (*dto.EventDto, error) {
	eventDto := &dto.EventDto{
		UserId:    eventModel.UserId,
		Date:      eventModel.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return eventDto, nil
}

func (c *ConverterEvents) MapEventModelToCreateEventDto(eventModel *model.EventModel) (*dto.CreateEventDto, error) {
	createEventDto := &dto.CreateEventDto{
		UserId: eventModel.UserId,
		Date:   eventModel.Date,
	}

	return createEventDto, nil
}

func (c *ConverterEvents) MapEventModelToUpdateEventDto(eventModel *model.EventModel) (*dto.UpdateEventDto, error) {
	updateEventDto := &dto.UpdateEventDto{
		UserId: eventModel.UserId,
		Date:   eventModel.Date,
	}

	return updateEventDto, nil
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
	eventResponse := &response.EventResponse{
		UserId:    eventDto.UserId,
		Date:      eventDto.Date,
		CreatedAt: eventDto.CreatedAt,
		UpdatedAt: eventDto.UpdatedAt,
	}

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
