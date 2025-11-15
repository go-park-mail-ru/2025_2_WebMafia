package postgres

import "database/sql"

type Repository struct {
	Conn *sql.DB
}

func NewTicketRepository(Conn *sql.DB) *Repository {
	return &Repository{Conn: Conn}
}
