package http

type CSRFManager interface {
	Generate(userID, sessionID string) (string, error)
}

type Handler struct {
	csrfManager CSRFManager
}

func NewHandler(csrfManager CSRFManager) *Handler {
	return &Handler{
		csrfManager: csrfManager,
	}
}
