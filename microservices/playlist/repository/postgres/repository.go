package postgres

import "database/sql"

type Repository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) *Repository {
	return &Repository{Conn: conn}
}
