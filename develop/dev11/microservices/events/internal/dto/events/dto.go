package events

import "time"

type EventDto struct {
	UserId    string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventsDto = []EventDto

type EventPeriodDto struct {
	From time.Time
	To   time.Time
}

type EventsForDayDto struct {
	UserId string
	Date   time.Time
}

type EventsForWeekDto struct {
	UserId string
	Date   time.Time
}

type EventsForMonthDto struct {
	UserId string
	Date   time.Time
}

type CreateEventDto struct {
	UserId string
	Date   time.Time
}

type UpdateEventDto struct {
	UserId string
	Date   time.Time
}

func NewEventDto(
	userId string,
	date time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *EventDto {
	return &EventDto{
		UserId:    userId,
		Date:      date,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewEventsForDayDto(
	userId string,
	date time.Time,
) *EventsForDayDto {
	return &EventsForDayDto{
		UserId: userId,
		Date:   date,
	}
}

func NewEventPeriodDto(
	from time.Time,
	to time.Time,
) *EventPeriodDto {
	return &EventPeriodDto{
		From: from,
		To:   to,
	}
}

func NewEventsForWeekDto(
	userId string,
	date time.Time,
) *EventsForWeekDto {
	return &EventsForWeekDto{
		UserId: userId,
		Date:   date,
	}
}

func NewEventsForMonthDto(
	userId string,
	date time.Time,
) *EventsForMonthDto {
	return &EventsForMonthDto{
		UserId: userId,
		Date:   date,
	}
}

func NewCreateEventDto(
	userId string,
	date time.Time,
) *CreateEventDto {
	return &CreateEventDto{
		UserId: userId,
		Date:   date,
	}
}

func NewUpdateEventDto(
	userId string,
	date time.Time,
) *UpdateEventDto {
	return &UpdateEventDto{
		UserId: userId,
		Date:   date,
	}
}
