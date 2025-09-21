package response

import (
	"bytes"
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

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Printf("could not encode response to json: %v", err)
		http.Error(w, "server error: could not encode response to json", http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
