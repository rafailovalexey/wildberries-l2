package tests

import (
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
	providerEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/provider/events"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	response "github.com/emptyhopes/wildberries-l2-dev11/internal/response/events"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapCreateEventRequestToCreateEventDto(t *testing.T) {
	converter := NewConverterEvents()

	createEventRequest := request.NewCreateEventRequest(
		"123",
		"2023-12-04",
	)

	createEventDto := dto.NewCreateEventDto(
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
	)

	result, err := converter.MapCreateEventRequestToCreateEventDto(createEventRequest)

	assert.Nil(t, err)
	assert.Equal(t, createEventDto, result)
}

func TestMapUpdateEventRequestToUpdateEventDto(t *testing.T) {
	converter := NewConverterEvents()

	updateEventRequest := request.NewUpdateEventRequest(
		1,
		"123",
		"2023-12-04",
	)

	updateEventDto := dto.NewUpdateEventDto(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
	)

	result, err := converter.MapUpdateEventRequestToUpdateEventDto(updateEventRequest)

	assert.Nil(t, err)
	assert.Equal(t, updateEventDto, result)
}

func TestMapEventModelToEventDto(t *testing.T) {
	converter := NewConverterEvents()

	eventModel := model.NewEventModel(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventDto := dto.NewEventDto(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
	)

	result, err := converter.MapEventModelToEventDto(eventModel)

	assert.Nil(t, err)
	assert.Equal(t, eventDto, result)
}

func TestMapEventsModelToEventsDto(t *testing.T) {
	converter := NewConverterEvents()

	eventModel1 := model.NewEventModel(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventModel2 := model.NewEventModel(
		2,
		"456",
		time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventsModel := &model.EventsModel{*eventModel1, *eventModel2}

	eventDto1 := dto.NewEventDto(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		eventModel1.CreatedAt,
		eventModel1.UpdatedAt,
	)

	eventDto2 := dto.NewEventDto(
		2,
		"456",
		time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		eventModel2.CreatedAt,
		eventModel2.UpdatedAt,
	)

	eventsDto := &dto.EventsDto{*eventDto1, *eventDto2}

	result, err := converter.MapEventsModelToEventsDto(eventsModel)

	assert.Nil(t, err)
	assert.Equal(t, eventsDto, result)
}

func TestMapEventDtoToEventResponse(t *testing.T) {
	converter := NewConverterEvents()

	eventDto := dto.NewEventDto(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventResponse := response.NewEventResponse(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		eventDto.CreatedAt,
		eventDto.UpdatedAt,
	)

	result, err := converter.MapEventDtoToEventResponse(eventDto)

	assert.Nil(t, err)
	assert.Equal(t, eventResponse, result)
}

func TestMapEventsDtoToEventsResponse(t *testing.T) {
	converter := NewConverterEvents()

	eventDto1 := dto.NewEventDto(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventDto2 := dto.NewEventDto(
		2,
		"456",
		time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now(),
	)

	eventsDto := &dto.EventsDto{*eventDto1, *eventDto2}

	eventResponse1 := response.NewEventResponse(
		1,
		"123",
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		eventDto1.CreatedAt,
		eventDto1.UpdatedAt,
	)

	eventResponse2 := response.NewEventResponse(
		2,
		"456",
		time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		eventDto2.CreatedAt,
		eventDto2.UpdatedAt,
	)

	eventsResponse := &response.EventsResponse{*eventResponse1, *eventResponse2}

	result, err := converter.MapEventsDtoToEventsResponse(eventsDto)

	assert.Nil(t, err)
	assert.Equal(t, eventsResponse, result)
}

func NewConverterEvents() converter.ConverterEventsInterface {
	provider := providerEvents.NewProviderEvents()

	c := provider.GetConverterEvents()

	return c
}
