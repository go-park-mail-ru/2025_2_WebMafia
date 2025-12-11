package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"spotify/microservices/catalog/dto"
	service_mock "spotify/microservices/catalog/mocks/service"
	"spotify/microservices/catalog/service"
)

func TestHandler_GetTrackComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service_mock.NewMockIService(ctrl)
	handler := NewHandler(mockSvc, nil, nil)

	trackID := uuid.New()
	comment := dto.Comment{
		ID:        uuid.New().String(),
		Text:      "Nice song!",
		UserLogin: "testuser",
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockSvc.EXPECT().
			GetCommentsByTrackID(gomock.Any(), trackID, uint64(100), uint64(0)).
			Return([]dto.Comment{comment}, nil)

		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackID.String()+"/comments", nil)
		req = mux.SetURLVars(req, map[string]string{"id": trackID.String()})
		rr := httptest.NewRecorder()

		handler.GetTrackComments(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Nice song!")
	})

	t.Run("not found (empty list)", func(t *testing.T) {
		mockSvc.EXPECT().
			GetCommentsByTrackID(gomock.Any(), trackID, uint64(100), uint64(0)).
			Return(nil, service.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackID.String()+"/comments", nil)
		req = mux.SetURLVars(req, map[string]string{"id": trackID.String()})
		rr := httptest.NewRecorder()

		handler.GetTrackComments(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, "[]", rr.Body.String())
	})

	t.Run("invalid track id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks/invalid/comments", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		rr := httptest.NewRecorder()

		handler.GetTrackComments(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		mockSvc.EXPECT().
			GetCommentsByTrackID(gomock.Any(), trackID, uint64(100), uint64(0)).
			Return(nil, errors.New("db error"))

		req := httptest.NewRequest(http.MethodGet, "/tracks/"+trackID.String()+"/comments", nil)
		req = mux.SetURLVars(req, map[string]string{"id": trackID.String()})
		rr := httptest.NewRecorder()

		handler.GetTrackComments(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tracks//comments", nil)
		rr := httptest.NewRecorder()

		handler.GetTrackComments(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
