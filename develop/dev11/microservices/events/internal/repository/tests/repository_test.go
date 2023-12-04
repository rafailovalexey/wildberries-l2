package tests

import (
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	mockRepository "github.com/emptyhopes/wildberries-l2-dev11/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRepositoryCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repository := mockRepository.NewMockRepositoryEventsInterface(ctrl)

	createEventDto := dto.NewCreateEventDto(
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
	)

	eventDto := dto.NewEventDto(
		1,
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
		time.Now(),
		time.Now(),
	)

	repository.EXPECT().CreateEvent(createEventDto).Return(eventDto, nil)
	result, err := repository.CreateEvent(createEventDto)

	assert.Nil(t, err)
	assert.Equal(t, result, eventDto)
}

func TestRepositoryUpdateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repository := mockRepository.NewMockRepositoryEventsInterface(ctrl)

	updateEventDto := dto.NewUpdateEventDto(
		1,
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
	)

	eventDto := dto.NewEventDto(
		1,
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
		time.Now(),
		time.Now(),
	)

	repository.EXPECT().UpdateEvent(updateEventDto).Return(eventDto, nil)
	result, err := repository.UpdateEvent(updateEventDto)

	assert.Nil(t, err)
	assert.Equal(t, result, eventDto)
}

func TestRepositoryGetEventsByUserIdAndPeriod(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repository := mockRepository.NewMockRepositoryEventsInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"

	period := dto.NewEventPeriodDto(
		time.Now(),
		time.Now(),
	)

	eventDto1 := dto.NewEventDto(
		1,
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
		time.Now(),
		time.Now(),
	)

	eventDto2 := dto.NewEventDto(
		2,
		"4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469",
		time.Now(),
		time.Now(),
		time.Now(),
	)

	eventsDto := &dto.EventsDto{*eventDto1, *eventDto2}

	repository.EXPECT().GetEventsByUserIdAndPeriod(id, period).Return(eventsDto, nil)
	result, err := repository.GetEventsByUserIdAndPeriod(id, period)

	assert.Nil(t, err)
	assert.Equal(t, result, eventsDto)
}
