package dto

import "time"

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type UpdateTicketRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Rating      *int    `json:"rating,omitempty"`
}

type TicketResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	Rating      *int32    `json:"rating,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type StatisticsResponse struct {
	TotalTickets   int            `json:"total_tickets"`
	ByStatus       map[string]int `json:"by_status"`
	ByCategory     map[string]int `json:"by_category"`
}
