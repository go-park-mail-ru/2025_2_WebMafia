package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(protected, csrfProtected, adminProtected, adminCsrfProtected *mux.Router) {
	csrfProtected.HandleFunc("/support/tickets", h.CreateTicket).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/support/tickets", h.GetUserTickets).Methods(http.MethodGet, http.MethodOptions)
	csrfProtected.HandleFunc("/support/tickets/{id}", h.UpdateTicket).Methods(http.MethodPut, http.MethodOptions)

	adminProtected.HandleFunc("/support/tickets", h.GetAllTickets).Methods(http.MethodGet, http.MethodOptions)
	adminProtected.HandleFunc("/support/statistics", h.GetStatistics).Methods(http.MethodGet, http.MethodOptions)
	adminCsrfProtected.HandleFunc("/support/tickets/{id}/status", h.UpdateTicketStatusByAdmin).Methods(http.MethodPut, http.MethodOptions)
}
