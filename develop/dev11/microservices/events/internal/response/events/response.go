package events

import "time"

type EventResponse struct {
	Id        int64     `json:"id"`
	UserId    string    `json:"user_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventsResponse = []EventResponse

func NewEventResponse(
	id int64,
	userId string,
	date time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *EventResponse {
	return &EventResponse{
		Id:        id,
		UserId:    userId,
		Date:      date,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
