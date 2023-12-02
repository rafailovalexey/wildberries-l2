package events

import "time"

type EventResponse struct {
	UserId    string    `json:"user_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventsResponse = []EventResponse
