package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(protected, csrfProtected *mux.Router) {
	csrfProtected.HandleFunc("/support/tickets", h.CreateTicket).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/support/tickets", h.GetUserTickets).Methods(http.MethodGet, http.MethodOptions)
	csrfProtected.HandleFunc("/support/tickets/{id}", h.UpdateTicket).Methods(http.MethodPut, http.MethodOptions)
	protected.HandleFunc("/support/tickets/all", h.GetAllTickets).Methods(http.MethodGet, http.MethodOptions)
}
