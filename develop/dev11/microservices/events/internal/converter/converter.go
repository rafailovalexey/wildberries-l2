package converter

import (
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
	request "github.com/emptyhopes/wildberries-l2-dev11/internal/request/events"
	response "github.com/emptyhopes/wildberries-l2-dev11/internal/response/events"
)

type ConverterEventsInterface interface {
	MapCreateEventRequestToCreateEventDto(*request.CreateEventRequest) (*dto.CreateEventDto, error)
	MapUpdateEventRequestToUpdateEventDto(*request.UpdateEventRequest) (*dto.UpdateEventDto, error)

	MapEventDtoToEventModel(*dto.EventDto) (*model.EventModel, error)
	MapEventsDtoToEventsModel(*dto.EventsDto) (*model.EventsModel, error)
	MapCreateEventDtoToEventModel(*dto.CreateEventDto) (*model.EventModel, error)
	MapUpdateEventDtoToEventModel(*dto.UpdateEventDto) (*model.EventModel, error)

	MapEventModelToEventDto(*model.EventModel) (*dto.EventDto, error)
	MapEventsModelToEventsDto(*model.EventsModel) (*dto.EventsDto, error)
	MapEventModelToCreateEventDto(*model.EventModel) (*dto.CreateEventDto, error)
	MapEventModelToUpdateEventDto(*model.EventModel) (*dto.UpdateEventDto, error)

	MapEventDtoToEventResponse(*dto.EventDto) (*response.EventResponse, error)
	MapEventsDtoToEventsResponse(*dto.EventsDto) (*response.EventsResponse, error)
}
