package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"spotify/internal/model"
	"spotify/pkg/response"
)

type AlbumsResponse struct {
	Albums []model.Album `json:"albums"`
}

type AlbumResponse struct {
	Album model.Album `json:"album"`
}

func (h *Handlers) GetAllAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albums, _ := h.store.GetAllAlbums()
	for i := range albums {
		enrichAlbumURLs(h.cfg, r, &albums[i])
	}
	response.JSON(w, http.StatusOK, AlbumsResponse{Albums: albums})
}

func (h *Handlers) GetAlbumByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumID := vars["id"]

	albums, _ := h.store.GetAllAlbums()
	for _, album := range albums {
		if fmt.Sprint(album.AlbumID) == albumID {
			enrichAlbumURLs(h.cfg, r, &album)
			response.JSON(w, http.StatusOK, AlbumResponse{Album: album})
			return
		}
	}
	response.JSON(w, http.StatusNotFound, response.ErrorResponse{Error: "album not found"})

}
