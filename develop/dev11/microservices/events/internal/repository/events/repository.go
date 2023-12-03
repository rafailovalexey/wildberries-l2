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

	eventModel, err := insertEvent(pool, createEventDto)

	if err != nil {
		return nil, err
	}

	eventDto, err := r.converterEvents.MapEventModelToEventDto(eventModel)

	if err != nil {
		return nil, err
	}

	return eventDto, nil
}

func insertEvent(
	pool *pgxpool.Pool,
	createEventDto *dto.CreateEventDto,
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
       RETURNING id, user_id, date, created_at, updated_at;
   `

	eventModel := &model.EventModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		createEventDto.UserId,
		createEventDto.Date,
		time.Now(),
		time.Now(),
	).Scan(
		&eventModel.Id,
		&eventModel.UserId,
		&eventModel.Date,
		&eventModel.CreatedAt,
		&eventModel.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return eventModel, nil
}

func (r *RepositoryEvents) UpdateEvent(updateEventDto *dto.UpdateEventDto) (*dto.EventDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	eventModel, err := updateEvent(pool, updateEventDto)

	if err != nil {
		return nil, err
	}

	eventDto, err := r.converterEvents.MapEventModelToEventDto(eventModel)

	if err != nil {
		return nil, err
	}

	return eventDto, nil
}

func updateEvent(
	pool *pgxpool.Pool,
	updateEventDto *dto.UpdateEventDto,
) (*model.EventModel, error) {
	query := `
       UPDATE events
		SET
			user_id = $1,
		    date = $2,
		    updated_at = $3
       WHERE id = $4
       RETURNING id, user_id, date, created_at, updated_at;
   `

	eventModel := &model.EventModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		updateEventDto.UserId,
		updateEventDto.Date,
		time.Now(),
		updateEventDto.Id,
	).Scan(
		&eventModel.Id,
		&eventModel.UserId,
		&eventModel.Date,
		&eventModel.CreatedAt,
		&eventModel.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return eventModel, nil
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

	eventsDto, err := r.converterEvents.MapEventsModelToEventsDto(eventsModel)

	if err != nil {
		return nil, err
	}

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
			&eventModel.Id,
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
