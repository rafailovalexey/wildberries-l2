package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

type DatabaseInterface interface {
	Initialize()
	GetPool() *pgxpool.Pool
	CreateTables(*pgxpool.Pool)
}

type Database struct {
	credentials string
}

var _ DatabaseInterface = (*Database)(nil)

func NewDatabase() *Database {
	db := &Database{}

	db.Initialize()

	pool := db.GetPool()
	defer pool.Close()

	db.CreateTables(pool)

	return db
}

func (d *Database) Initialize() {
	username := os.Getenv("POSTGRESQL_USERNAME")

	if username == "" {
		log.Panicf("specify the database user")
	}

	password := os.Getenv("POSTGRESQL_PASSWORD")

	if password == "" {
		log.Panicf("specify the database user password")
	}

	db := os.Getenv("POSTGRESQL_DATABASE")

	if db == "" {
		log.Panicf("indicate the name of the database")
	}

	hostname := os.Getenv("POSTGRESQL_HOSTNAME")

	if hostname == "" {
		log.Panicf("specify the database hostname")
	}

	port := os.Getenv("POSTGRESQL_PORT")

	if port == "" {
		log.Panicf("specify the database port")
	}

	sslmode := os.Getenv("POSTGRESQL_SSLMODE")

	if sslmode == "" {
		log.Panicf("specify the ssl mode of the database")
	}

	d.credentials = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, db, hostname, port, sslmode)
}

func (d *Database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Panicf("database connection error %v\n", err)
	}

	return pool
}

func (d *Database) CreateTables(pool *pgxpool.Pool) {
	createEmployeeTable(pool)
}

func createEmployeeTable(pool *pgxpool.Pool) {
	query := `
		CREATE TABLE IF NOT EXISTS events (
		    id SERIAL PRIMARY KEY,
			user_id UUID NOT NULL,
			date TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
		);
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Panicf("table creation error %v\n", err)
	}
}
