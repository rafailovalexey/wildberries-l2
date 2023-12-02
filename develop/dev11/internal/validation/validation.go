package validation

type ValidationEventsInterface interface {
	EventsForDayValidation(userId string, date string) error
	EventsForWeekValidation(userId string, date string) error
	EventsForMonthValidation(userId string, date string) error
}
