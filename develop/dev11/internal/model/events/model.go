package events

import "time"

type EventModel struct {
	Id        int64
	UserId    string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventsModel = []EventModel

func NewEventModel(
	id int64,
	userId string,
	date time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *EventModel {
	return &EventModel{
		Id:        id,
		UserId:    userId,
		Date:      date,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
