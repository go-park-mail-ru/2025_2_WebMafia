package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not encode response to json: %v", err)
		JSON(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

// 400 Bad Request
func BadRequestJSON(w http.ResponseWriter) {
	JSON(w, http.StatusBadRequest, ErrBadRequest)
}

// 401 Unauthorized
func UnauthorizedJSON(w http.ResponseWriter) {
	JSON(w, http.StatusUnauthorized, ErrUnauthorized)
}

// 403 Forbidden
func ForbiddenJSON(w http.ResponseWriter) {
	JSON(w, http.StatusForbidden, ErrForbidden)
}

// 404 Not Found
func NotFoundJSON(w http.ResponseWriter) {
	JSON(w, http.StatusNotFound, ErrNotFound)
}

// 409 Conflict
func ConflictJSON(w http.ResponseWriter) {
	JSON(w, http.StatusConflict, ErrConflict)
}

// 500 Internal Server Error
func InternalErrorJSON(w http.ResponseWriter) {
	JSON(w, http.StatusInternalServerError, ErrInternalServer)
}
