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

	if err, ok := data.(error); ok {
		log.Printf("ERROR: status %d: %v", statusCode, err)

		switch {
		case statusCode >= 400 && statusCode < 500:
			data = ErrBadRequest
			if statusCode == http.StatusUnauthorized {
				data = ErrUnauthorized
			}
			if statusCode == http.StatusNotFound {
				data = ErrNotFound
			}
			if statusCode == http.StatusConflict {
				data = ErrConflict
			}
		case statusCode >= 500:
			data = ErrInternalServer
		default:
			data = ErrorResponse{Error: http.StatusText(statusCode)}
		}
	}
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
