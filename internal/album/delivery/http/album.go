package http

import (
	"errors"
	"log"
	"net/http"
	"spotify/internal/album/dto"
	"spotify/internal/album/service"
	"spotify/pkg/response"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AlbumsResponse struct {
	Albums []dto.Album `json:"albums"`
}

type AlbumResponse struct {
	Album *dto.Album `json:"album"`
}

func (h *Handler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		log.Println("ERROR: delivery.GetAlbumByID: id is missing")
		response.BadRequestJSON(w)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("ERROR: delivery.GetAlbumByID: failed to parse id: %v", err)
		response.BadRequestJSON(w)
		return
	}

	album, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		log.Printf("ERROR: delivery.GetAlbumByID: service error: %v", err)
		if errors.Is(err, service.ErrNotFound) {
			response.NotFoundJSON(w)
			return
		}
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, AlbumResponse{Album: album})
}

func (h *Handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := h.service.GetAllAlbums(r.Context())
	if err != nil {
		log.Printf("ERROR: delivery.GetAllAlbums: service error: %v", err)
		response.InternalErrorJSON(w)
		return
	}

	response.JSON(w, http.StatusOK, AlbumsResponse{Albums: albums})
}
