package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"spotify/internal/track/dto"
	"spotify/internal/track/service"
	"spotify/pkg/response"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TracksResponse struct {
	Tracks []dto.Track `json:"tracks"`
}

type TrackResponse struct {
	Track *dto.Track `json:"track"`
}

func (h *Handler) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("ERROR: handler.GetTrackByID: failed to parse id: %v", err)
		response.BadRequestJSON(w)
		return
	}

	track, err := h.service.GetTrackByID(r.Context(), id)
	if err != nil {
		log.Printf("ERROR: handler.GetTrackByID: service error: %v", err)
		if errors.Is(err, service.ErrNotFound) {
			response.NotFoundJSON(w)
			return
		}
		response.InternalErrorJSON(w)
		return
	}
	response.JSON(w, http.StatusOK, TrackResponse{Track: track})
}

func (h *Handler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := h.service.GetAllTracks(r.Context())
	if err != nil {
		log.Printf("ERROR: handler.GetAllTracks: service error: %v", err)
		response.InternalErrorJSON(w)
		return
	}
	response.JSON(w, http.StatusOK, TracksResponse{Tracks: tracks})
}

func (h *Handler) getTracksByParam(w http.ResponseWriter, r *http.Request, paramName string, serviceCall func(ctx context.Context, id uuid.UUID) ([]dto.Track, error)) {
	id, err := uuid.Parse(mux.Vars(r)[paramName])
	if err != nil {
		log.Printf("ERROR: handler.getTracksByParam (%s): failed to parse id: %v", paramName, err)
		response.BadRequestJSON(w)
		return
	}

	tracks, err := serviceCall(r.Context(), id)
	if err != nil {
		log.Printf("ERROR: handler.getTracksByParam (%s): service error: %v", paramName, err)
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, TracksResponse{Tracks: tracks})
}

func (h *Handler) GetTracksByArtist(w http.ResponseWriter, r *http.Request) {
	h.getTracksByParam(w, r, "artistId", h.service.GetTracksByArtistID)
}

func (h *Handler) GetTracksByAlbum(w http.ResponseWriter, r *http.Request) {
	h.getTracksByParam(w, r, "albumId", h.service.GetTracksByAlbumID)
}

func (h *Handler) GetTracksByGenre(w http.ResponseWriter, r *http.Request) {
	h.getTracksByParam(w, r, "genreId", h.service.GetTracksByGenreID)
}
