package events

type CreateEventRequest struct {
	UserId string `json:"user_id"`
	Date   string `json:"date"`
}

type UpdateEventRequest struct {
	Id     int64  `json:"id,int64"`
	UserId string `json:"user_id"`
	Date   string `json:"date"`
}

func NewCreateEventRequest(
	userId string,
	date string,
) *CreateEventRequest {
	return &CreateEventRequest{
		UserId: userId,
		Date:   date,
	}
}

func NewUpdateEventRequest(
	id int64,
	userId string,
	date string,
) *UpdateEventRequest {
	return &UpdateEventRequest{
		Id:     id,
		UserId: userId,
		Date:   date,
	}
}
