package events

type CreateEventRequest struct {
	UserId string `json:"user_id"`
	Date   string `json:"date"`
}

type UpdateEventRequest struct {
	UserId string `json:"user_id"`
	Date   string `json:"date"`
}
