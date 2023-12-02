package events

import "time"

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

func NewEventsForDayDto(
	userId string,
	date time.Time,
) *EventsForDayDto {
	return &EventsForDayDto{
		UserId: userId,
		Date:   date,
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
