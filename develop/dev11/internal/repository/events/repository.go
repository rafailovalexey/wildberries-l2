package events

import (
	"context"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/converter"
	dto "github.com/emptyhopes/wildberries-l2-dev11/internal/dto/events"
	model "github.com/emptyhopes/wildberries-l2-dev11/internal/model/events"
	definition "github.com/emptyhopes/wildberries-l2-dev11/internal/repository"
	"github.com/emptyhopes/wildberries-l2-dev11/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"time"
)

type RepositoryEvents struct {
	converterEvents converter.ConverterEventsInterface
	database        storage.DatabaseInterface
	rwmutex         sync.RWMutex
}

var _ definition.RepositoryEventsInterface = (*RepositoryEvents)(nil)

func NewRepositoryEvents(
	converterEvents converter.ConverterEventsInterface,
	database storage.DatabaseInterface,
) *RepositoryEvents {
	return &RepositoryEvents{
		converterEvents: converterEvents,
		database:        database,
		rwmutex:         sync.RWMutex{},
	}
}

func (r *RepositoryEvents) CreateEvent(
	createEventDto *dto.CreateEventDto,
) (*dto.EventDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	eventModel := r.converterEvents.MapCreateEventDtoToEventModel(createEventDto)

	eventModel, err := insertEvent(pool, eventModel)

	if err != nil {
		return nil, err
	}

	eventDto := r.converterEvents.MapEventModelToEventDto(eventModel)

	return eventDto, nil
}

func insertEvent(
	pool *pgxpool.Pool,
	eventModel *model.EventModel,
) (*model.EventModel, error) {
	query := `
       INSERT INTO events (
			user_id,
			date,
			created_at,
			updated_at
       )
       VALUES (
           	$1,
           	$2,
           	$3,
            $4
       )
       RETURNING (
			id,
			user_id,
			date,
			created_at,
			updated_at
       );
   `

	createdEvent := &model.EventModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		eventModel.UserId,
		eventModel.Date,
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
	).Scan(
		&createdEvent.UserId,
		&createdEvent.Date,
		&createdEvent.CreatedAt,
		&createdEvent.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdEvent, nil
}

func (r *RepositoryEvents) UpdateEvent(updateEventDto *dto.UpdateEventDto) (*dto.EventDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	eventModel := r.converterEvents.MapUpdateEventDtoToEventModel(updateEventDto)

	eventModel, err := updateEvent(pool, eventModel)

	if err != nil {
		return nil, err
	}

	eventDto := r.converterEvents.MapEventModelToEventDto(eventModel)

	return eventDto, nil
}

func updateEvent(
	pool *pgxpool.Pool,
	eventModel *model.EventModel,
) (*model.EventModel, error) {
	query := `
       UPDATE events
		SET
		    date = $1,
		    updated_at = $2
       WHERE user_id = $3
       RETURNING (
			id,
			user_id,
			date,
			created_at,
			updated_at
       );
   `

	updatedEvent := &model.EventModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		eventModel.Date,
		time.Now(),
		eventModel.UserId,
	).Scan(
		&updatedEvent.UserId,
		&updatedEvent.Date,
		&updatedEvent.CreatedAt,
		&updatedEvent.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return updatedEvent, nil
}

func (r *RepositoryEvents) GetEventsByUserIdAndPeriod(
	userId string,
	eventPeriodDto *dto.EventPeriodDto,
) (*dto.EventsDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	eventsModel, err := getEventsByUserIdAndPeriod(pool, userId, eventPeriodDto)

	if err != nil {
		return nil, err
	}

	eventsDto := r.converterEvents.MapEventsModelToEventsDto(eventsModel)

	return eventsDto, nil
}

func getEventsByUserIdAndPeriod(
	pool *pgxpool.Pool,
	userId string,
	eventPeriodDto *dto.EventPeriodDto,
) (*model.EventsModel, error) {
	query := `
       SELECT
           	id,
			user_id,
			date,
			created_at,
			updated_at
       FROM events
       WHERE user_id = $1 AND date BETWEEN $2 AND $3
   `

	rows, err := pool.Query(
		context.Background(),
		query,
		userId,
		eventPeriodDto.From,
		eventPeriodDto.To,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	eventsModel := make(model.EventsModel, 0, 10)

	for rows.Next() {
		eventModel := model.EventModel{}

		err = rows.Scan(
			&eventModel.UserId,
			&eventModel.Date,
			&eventModel.CreatedAt,
			&eventModel.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		eventsModel = append(eventsModel, eventModel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &eventsModel, nil
}
