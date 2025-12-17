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

func TestHandler_GeneratePlaylistMeta(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().
			GeneratePlaylistMeta(gomock.Any(), id).
			Return(&dto.GeneratedMeta{
				Title:       "T",
				Description: "D",
				Source:      "ai",
			}, nil)

		req := httptest.NewRequest("POST", "/playlists/"+id.String()+"/generate-meta", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		rr := httptest.NewRecorder()
		handler.GeneratePlaylistMeta(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid playlist id", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/playlists/bad/generate-meta", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})

		rr := httptest.NewRecorder()
		handler.GeneratePlaylistMeta(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc.EXPECT().
			GeneratePlaylistMeta(gomock.Any(), id).
			Return(nil, errors.New("err"))

		req := httptest.NewRequest("POST", "/playlists/"+id.String()+"/generate-meta", nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		rr := httptest.NewRecorder()
		handler.GeneratePlaylistMeta(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandler_ConfirmPlaylistMeta(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)
	id := uuid.New()

	body, _ := json.Marshal(dto.ConfirmGeneratedMetaRequest{
		Title:       "New title",
		Description: "New desc",
	})

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().
			ConfirmPlaylistMeta(gomock.Any(), id, "New title", "New desc").
			Return(nil)

		req := httptest.NewRequest(
			"POST",
			"/playlists/"+id.String()+"/confirm-meta",
			bytes.NewReader(body),
		)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		rr := httptest.NewRecorder()
		handler.ConfirmPlaylistMeta(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid playlist id", func(t *testing.T) {
		req := httptest.NewRequest(
			"POST",
			"/playlists/bad/confirm-meta",
			bytes.NewReader(body),
		)
		req = mux.SetURLVars(req, map[string]string{"id": "bad"})

		rr := httptest.NewRecorder()
		handler.ConfirmPlaylistMeta(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("invalid body", func(t *testing.T) {
		req := httptest.NewRequest(
			"POST",
			"/playlists/"+id.String()+"/confirm-meta",
			bytes.NewReader([]byte("bad json")),
		)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		rr := httptest.NewRecorder()
		handler.ConfirmPlaylistMeta(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("service error", func(t *testing.T) {
		mockSvc.EXPECT().
			ConfirmPlaylistMeta(gomock.Any(), id, "New title", "New desc").
			Return(errors.New("err"))

		req := httptest.NewRequest(
			"POST",
			"/playlists/"+id.String()+"/confirm-meta",
			bytes.NewReader(body),
		)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		rr := httptest.NewRecorder()
		handler.ConfirmPlaylistMeta(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func Test_parsePagination(t *testing.T) {
	t.Run("empty query", func(t *testing.T) {
		limit, offset := parsePagination(url.Values{})
		assert.Equal(t, uint64(defaultLimit), limit)
		assert.Equal(t, uint64(defaultOffset), offset)
	})

	t.Run("invalid values", func(t *testing.T) {
		q := url.Values{
			"limit":  []string{"bad"},
			"offset": []string{"bad"},
		}
		limit, offset := parsePagination(q)
		assert.Equal(t, uint64(defaultLimit), limit)
		assert.Equal(t, uint64(defaultOffset), offset)
	})

	t.Run("limit too large", func(t *testing.T) {
		q := url.Values{
			"limit": []string{"999999"},
		}
		limit, _ := parsePagination(q)
		assert.Equal(t, uint64(maxLimit), limit)
	})
}

func TestHandler_GetMyPlaylists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)

	uid := uuid.New()

	mockSvc.EXPECT().
		GetPlaylistsByUser(gomock.Any(), gomock.Any()).
		Return([]dto.Playlist{
			{ID: uuid.New().String()},
		}, nil)

	mockSvc.EXPECT().
		GetPlaylistWithTracks(gomock.Any(), gomock.Any()).
		Return(&dto.Playlist{}, nil)

	req := httptest.NewRequest("GET", "/playlists/my", nil)
	ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{
		UserID: uid.String(),
	})

	rr := httptest.NewRecorder()
	handler.GetMyPlaylists(rr, req.WithContext(ctx))

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_GetMyPlaylists_NoUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)

	req := httptest.NewRequest("GET", "/playlists/my", nil)
	rr := httptest.NewRecorder()

	handler.GetMyPlaylists(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestHandler_GetMyPlaylists_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil)

	uid := uuid.New()

	mockSvc.EXPECT().
		GetPlaylistsByUser(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("err"))

	req := httptest.NewRequest("GET", "/playlists/my", nil)
	ctx := middleware.ContextWithClaims(req.Context(), &jwtmanager.Claims{
		UserID: uid.String(),
	})

	rr := httptest.NewRecorder()
	handler.GetMyPlaylists(rr, req.WithContext(ctx))

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestValidatePlaylistAvatar(t *testing.T) {
	h := &Handler{
		allowedAvatarTypes: []string{"image/png"},
	}

	t.Run("empty file", func(t *testing.T) {
		err := h.validatePlaylistAvatar("image/png", 0)
		assert.Error(t, err)
	})

	t.Run("too large", func(t *testing.T) {
		err := h.validatePlaylistAvatar("image/png", maxPlaylistAvatarSize+1)
		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		err := h.validatePlaylistAvatar("image/jpeg", 10)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		err := h.validatePlaylistAvatar("image/png", 10)
		assert.NoError(t, err)
	})
}

func TestHandler_AddArtistToFavorite_NoUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := NewHandler(service_mock.NewMockIService(ctrl), nil)

	req := httptest.NewRequest("POST", "/artists/a1/favorite", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "a1"})

	rr := httptest.NewRecorder()
	handler.AddArtistToFavorite(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
