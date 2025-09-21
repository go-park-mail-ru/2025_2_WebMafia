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
		log.Printf("could not encode response to json: %v", err)
		JSON(w, http.StatusInternalServerError, ErrorResponse{Error: "server error: could not encode response to json"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
