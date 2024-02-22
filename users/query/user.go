package query

import (
	"context"
	"database/sql"
)

type QueryExecuter struct {
	db *sql.DB
}

func NewQueryExecutor(db *sql.DB) *QueryExecuter {
	return &QueryExecuter{db: db}
}

func (s *QueryExecuter) GetUser(ctx context.Context, id string) (*sql.Rows, error) {
	rows, err := s.db.Query(`SELECT id, username, password, email FROM users.users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *QueryExecuter) CreateUser(ctx context.Context, id string, username string, password string, email string) (*sql.Rows, error) {
	rows, err := s.db.Query("INSERT INTO users.users (id, username, password, email) VALUES ($1, $2, $3, $4) RETURNING id, username, password, email", id, username, password, email)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *QueryExecuter) UpdateUser(ctx context.Context, id string, username string, password string, email string) (*sql.Rows, error) {
	rows, err := s.db.Query("UPDATE users.users SET username = $2, password = $3, email = $4 WHERE id = $1 RETURNING id, username, password, email", id, username, password, email)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *QueryExecuter) DeleteUser(ctx context.Context, id string) error {
	_, err := s.db.Exec("DELETE FROM users.users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
