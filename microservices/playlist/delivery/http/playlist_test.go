package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"net/url"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/internal/middleware"
	"spotify/microservices/playlist/dto"
	"spotify/microservices/playlist/service"
	service_mock "spotify/mocks/playlist/service"
	"spotify/pkg/jwtmanager"
)

func TestHandler_CreatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	uid := uuid.New()
	body, _ := json.Marshal(dto.CreatePlaylistRequest{Title: "P1"})

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().CreatePlaylist(gomock.Any(), gomock.Any()).Return(&dto.Playlist{Title: "P1"}, nil)
		req := httptest.NewRequest("POST", "/playlists", bytes.NewReader(body))
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.CreatePlaylist(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/playlists", bytes.NewReader([]byte("bad json")))
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.CreatePlaylist(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("no user id", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/playlists", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		handler.CreatePlaylist(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_GetPlaylistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().GetPlaylistWithTracks(gomock.Any(), id).Return(&dto.Playlist{}, nil)
		req := httptest.NewRequest("GET", "/playlists/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.GetPlaylistByID(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/playlists/bad", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})
		rr := httptest.NewRecorder()
		handler.GetPlaylistByID(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("not found", func(t *testing.T) {
		mockSvc.EXPECT().GetPlaylistWithTracks(gomock.Any(), id).Return(nil, service.ErrNotFound)
		req := httptest.NewRequest("GET", "/playlists/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.GetPlaylistByID(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandler_UpdatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()
	body, _ := json.Marshal(dto.UpdatePlaylistRequest{})

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().UpdatePlaylist(gomock.Any(), gomock.Any()).Return(&dto.Playlist{}, nil)
		req := httptest.NewRequest("PUT", "/playlists/"+id.String(), bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.UpdatePlaylist(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/playlists/bad", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})
		rr := httptest.NewRecorder()
		handler.UpdatePlaylist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/playlists/"+id.String(), bytes.NewReader([]byte("bad")))
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.UpdatePlaylist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_DeletePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().DeletePlaylist(gomock.Any(), gomock.Any()).Return(nil)
		req := httptest.NewRequest("DELETE", "/playlists/"+id.String(), nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.DeletePlaylist(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/playlists/bad", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})
		rr := httptest.NewRecorder()
		handler.DeletePlaylist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_UploadPlaylistAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, []string{"image/jpeg"})
	id := uuid.New()
	uid := uuid.New()

	t.Run("success", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Disposition", `form-data; name="avatar"; filename="test.jpg"`)
		partHeader.Set("Content-Type", "image/jpeg")
		part, _ := writer.CreatePart(partHeader)
		io.Copy(part, strings.NewReader("data"))
		writer.Close()

		mockSvc.EXPECT().UploadPlaylistAvatar(gomock.Any(), gomock.Any()).Return(&dto.UploadPlaylistAvatarResponse{URL: "url"}, nil)

		req := httptest.NewRequest("POST", "/playlists/"+id.String()+"/avatar", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})

		rr := httptest.NewRecorder()
		handler.UploadPlaylistAvatar(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid file type", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Disposition", `form-data; name="avatar"; filename="test.exe"`)
		partHeader.Set("Content-Type", "application/exe")
		part, _ := writer.CreatePart(partHeader)
		io.Copy(part, strings.NewReader("data"))
		writer.Close()

		req := httptest.NewRequest("POST", "/playlists/"+id.String()+"/avatar", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})

		rr := httptest.NewRecorder()
		handler.UploadPlaylistAvatar(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("unauthorized", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/playlists/"+id.String()+"/avatar", nil)
		rr := httptest.NewRecorder()
		handler.UploadPlaylistAvatar(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestHandler_DeletePlaylistAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()
	uid := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().DeletePlaylistAvatar(gomock.Any(), gomock.Any()).Return(nil)
		req := httptest.NewRequest("DELETE", "/playlists/"+id.String()+"/avatar", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.DeletePlaylistAvatar(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc.EXPECT().DeletePlaylistAvatar(gomock.Any(), gomock.Any()).Return(errors.New("err"))
		req := httptest.NewRequest("DELETE", "/playlists/"+id.String()+"/avatar", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.DeletePlaylistAvatar(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_AddTrackToPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()
	body, _ := json.Marshal(dto.AddTrackToPlaylistRequest{TrackID: "t1"})

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().AddTrackToPlaylist(gomock.Any(), gomock.Any()).Return(nil)
		req := httptest.NewRequest("POST", "/pl", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.AddTrackToPlaylist(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/pl", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})
		rr := httptest.NewRecorder()
		handler.AddTrackToPlaylist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_RemoveTrackFromPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()
	body, _ := json.Marshal(dto.RemoveTrackFromPlaylistRequest{TrackID: "t1"})

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().RemoveTrackFromPlaylist(gomock.Any(), gomock.Any()).Return(nil)
		req := httptest.NewRequest("DELETE", "/pl", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
		rr := httptest.NewRecorder()
		handler.RemoveTrackFromPlaylist(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/pl", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})
		rr := httptest.NewRecorder()
		handler.RemoveTrackFromPlaylist(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_Favorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	uid := uuid.New()

	t.Run("AddTrackToFavorite no user", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/fav", nil)
		rr := httptest.NewRecorder()
		handler.AddTrackToFavorite(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("AddTrackToFavorite success", func(t *testing.T) {
		mockSvc.EXPECT().AddTrackToFavorite(gomock.Any(), gomock.Any()).Return(nil)
		body, _ := json.Marshal(dto.AddTrackToFavoriteRequest{TrackID: "t1"})
		req := httptest.NewRequest("POST", "/fav", bytes.NewReader(body))
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.AddTrackToFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("GetFavoritePlaylist success", func(t *testing.T) {
		mockSvc.EXPECT().GetFavoritePlaylist(gomock.Any(), gomock.Any()).Return(&dto.Playlist{ID: uuid.New().String()}, nil)
		mockSvc.EXPECT().GetPlaylistWithTracks(gomock.Any(), gomock.Any()).Return(&dto.Playlist{}, nil)

		req := httptest.NewRequest("GET", "/fav", nil)
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		handler.GetFavoritePlaylist(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_AddAlbumToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user id", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/albums/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rr := httptest.NewRecorder()
		h.AddAlbumToFavorite(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user id", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/albums/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.AddAlbumToFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().
			AddAlbumToFavorite(gomock.Any(), gomock.Any()).
			Return(nil)

		req := httptest.NewRequest("POST", "/albums/123/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "123"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.AddAlbumToFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_RemoveAlbumFromFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/albums/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rr := httptest.NewRecorder()
		h.RemoveAlbumFromFavorite(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user id", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/albums/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.RemoveAlbumFromFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().RemoveAlbumFromFavorite(gomock.Any(), gomock.Any()).Return(nil)

		req := httptest.NewRequest("DELETE", "/albums/123/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "123"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.RemoveAlbumFromFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_GetFavoriteAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/albums/fav", nil)
		rr := httptest.NewRecorder()
		h.GetFavoriteAlbums(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/albums/fav", nil)
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.GetFavoriteAlbums(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().
			GetFavoriteAlbums(gomock.Any(), uid).
			Return([]dto.FavoriteAlbum{
				{
					ID:        "alb1",
					Title:     "A1",
					CreatorID: uid.String(),
				},
			}, nil)

		req := httptest.NewRequest("GET", "/albums/fav", nil)
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.GetFavoriteAlbums(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_AddArtistToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/artists/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rr := httptest.NewRecorder()
		h.AddArtistToFavorite(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/artists/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.AddArtistToFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().AddArtistToFavorite(gomock.Any(), gomock.Any()).Return(nil)

		req := httptest.NewRequest("POST", "/artists/123/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "123"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.AddArtistToFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_RemoveArtistFromFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/artists/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		rr := httptest.NewRecorder()
		h.RemoveArtistFromFavorite(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/artists/1/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.RemoveArtistFromFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().RemoveArtistFromFavorite(gomock.Any(), gomock.Any()).Return(nil)

		req := httptest.NewRequest("DELETE", "/artists/123/fav", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "123"})
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.RemoveArtistFromFavorite(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHandler_GetFavoriteArtists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service_mock.NewMockIService(ctrl)
	h := NewHandler(svc, nil)
	uid := uuid.New()

	t.Run("no user", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/artists/fav", nil)
		rr := httptest.NewRecorder()
		h.GetFavoriteArtists(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("invalid user id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/artists/fav", nil)
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: "bad"})
		rr := httptest.NewRecorder()
		h.GetFavoriteArtists(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("success", func(t *testing.T) {
		svc.EXPECT().
			GetFavoriteArtists(gomock.Any(), uid).
			Return([]dto.FavoriteArtist{
				{
					ID:        "a1",
					Name:      "A1",
					CreatorID: uid.String(),
				},
			}, nil)

		req := httptest.NewRequest("GET", "/artists/fav", nil)
		ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{UserID: uid.String()})
		rr := httptest.NewRecorder()
		h.GetFavoriteArtists(rr, req.WithContext(ctx))
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestParsePagination(t *testing.T) {
	q := url.Values{}
	q.Set("limit", "2000")
	q.Set("offset", "10")

	limit, offset := parsePagination(q)
	assert.Equal(t, uint64(1000), limit)
	assert.Equal(t, uint64(10), offset)

	q = url.Values{}
	limit, offset = parsePagination(q)
	assert.Equal(t, uint64(100), limit)
	assert.Equal(t, uint64(0), offset)
}
