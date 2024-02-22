package sampleDatabase

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

func NewConnectionServer(db *sql.DB) *Database {
	return &Database{
		Db: db,
	}
}

func (d *Database) Create(host, port, user, password, name string) (*sql.DB, error) {
	dbHost := host
	dbPort := port
	dbUser := user
	dbPassword := password
	dbName := name

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	d.Db = db

	return d.Db, nil
}

func (d *Database) Get() (*sql.DB, error) {
	if d.Db == nil {
		return nil, errors.New("database connection is nil")
	}
	return d.Db, nil
}

func (d *Database) Close() error {
	return d.Db.Close()
}
