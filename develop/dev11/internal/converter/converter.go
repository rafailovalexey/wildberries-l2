package converter

import (
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
)

type ConverterEventsInterface interface {
	MapEventDtoToEventModel(*dto.EventDto) *model.EventModel
	MapEventsDtoToEventsModel(*dto.EventsDto) *model.EventsModel
	MapCreateEventDtoToEventModel(*dto.CreateEventDto) *model.EventModel
	MapUpdateEventDtoToEventModel(*dto.UpdateEventDto) *model.EventModel

	MapEventModelToEventDto(*model.EventModel) *dto.EventDto
	MapEventsModelToEventsDto(*model.EventsModel) *dto.EventsDto
	MapEventModelToCreateEventDto(*model.EventModel) *dto.CreateEventDto
	MapEventModelToUpdateEventDto(*model.EventModel) *dto.UpdateEventDto
}
