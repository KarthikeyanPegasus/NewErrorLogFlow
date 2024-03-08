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

func (q *QueryExecuter) CreatePost(ctx context.Context, id, content, parent string) (*sql.Rows, error) {
	rows, err := q.db.Query("INSERT INTO posts.posts (id, content, parent) VALUES ($1, $2, $3) RETURNING id, content, parent", id, content, parent)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (q *QueryExecuter) GetPost(ctx context.Context, id string) (*sql.Rows, error) {
	rows, err := q.db.Query(`SELECT id, content, parent FROM posts.posts WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (q *QueryExecuter) UpdatePost(ctx context.Context, id, content, parent string) (*sql.Rows, error) {
	rows, err := q.db.Query("UPDATE posts.posts SET content = $2, parent = $3 WHERE id = $1 RETURNING id, content, parent", id, content, parent)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (q *QueryExecuter) DeletePost(ctx context.Context, id string) (*sql.Rows, error) {
	rows, err := q.db.Query("DELETE FROM posts.posts WHERE id = $1 RETURNING id, content, parent", id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (q *QueryExecuter) DeleteAllPostOfUser(ctx context.Context, parent string) (*sql.Rows, error) {
	rows, err := q.db.Query("DELETE FROM posts.posts WHERE parent = $1 RETURNING id, content, parent", parent)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
