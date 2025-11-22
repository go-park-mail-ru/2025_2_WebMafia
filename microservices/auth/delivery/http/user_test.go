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
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"spotify/internal/middleware"
	mock_user "spotify/internal/mocks/user"
	"spotify/internal/user/dto"
	"spotify/internal/user/service"
	"spotify/pkg/jwtmanager"
)

func TestHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_user.NewMockIService(ctrl)
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	handler := NewHandler(mockUserService, jwtManager, nil, nil)

	loginReqPayload := loginRequest{
		Login:    "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(loginReqPayload)

	t.Run("success - 200 OK", func(t *testing.T) {
		serviceResponse := &dto.LoginResponse{ID: "some-uuid"}
		mockUserService.EXPECT().
			Login(gomock.Any(), dto.LoginRequest{Login: loginReqPayload.Login, Password: loginReqPayload.Password}).
			Return(serviceResponse, nil)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respBody loginResponse
		err := json.Unmarshal(rr.Body.Bytes(), &respBody)
		require.NoError(t, err)
		assert.Equal(t, serviceResponse.ID, respBody.ID)

		cookie := rr.Result().Cookies()[0]
		assert.Equal(t, sessionTokenCookie, cookie.Name)
		assert.NotEmpty(t, cookie.Value)
	})

	t.Run("fail - service returns validation error - 400 Bad Request", func(t *testing.T) {
		mockUserService.EXPECT().
			Login(gomock.Any(), gomock.Any()).
			Return(nil, service.ErrValidation)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("fail - invalid json - 400 Bad Request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte("{invalid json")))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("fail - request validation empty login", func(t *testing.T) {
		invalidPayload := loginRequest{Login: "", Password: "password123"}
		invalidBody, _ := json.Marshal(invalidPayload)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(invalidBody))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_user.NewMockIService(ctrl)
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	handler := NewHandler(mockUserService, jwtManager, nil, nil)

	registerReqPayload := registerRequest{
		Login:    "testuser123",
		Password: "password123",
		Email:    "test@test.com",
	}
	body, _ := json.Marshal(registerReqPayload)

	t.Run("success - 201 Created", func(t *testing.T) {
		serviceResponse := &dto.RegisterResponse{
			ID:    "some-uuid",
			Login: registerReqPayload.Login,
			Email: registerReqPayload.Email,
		}
		mockUserService.EXPECT().
			Register(gomock.Any(), gomock.Any()).
			Return(serviceResponse, nil)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
		rr := httptest.NewRecorder()

		handler.Register(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var respBody registerResponse
		err := json.Unmarshal(rr.Body.Bytes(), &respBody)
		require.NoError(t, err)
		assert.Equal(t, serviceResponse.ID, respBody.ID)
		assert.NotEmpty(t, rr.Result().Cookies())
	})

	t.Run("fail - validation error - 400 Bad Request", func(t *testing.T) {
		invalidPayload := registerRequest{Login: "sho"}
		invalidBody, _ := json.Marshal(invalidPayload)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(invalidBody))
		rr := httptest.NewRecorder()

		handler.Register(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("fail - service conflict error", func(t *testing.T) {
		mockUserService.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil, service.ErrConflict)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
		rr := httptest.NewRecorder()

		handler.Register(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})
}

func TestHandler_Logout(t *testing.T) {
	handler := NewHandler(nil, nil, nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	rr := httptest.NewRecorder()

	handler.Logout(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	cookie := rr.Result().Cookies()[0]
	assert.True(t, cookie.Expires.Before(time.Now()))
}

func TestHandler_Avatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_user.NewMockIService(ctrl)
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	authMW := middleware.NewAuthMiddleware(jwtManager)
	handler := NewHandler(mockUserService, jwtManager, nil, []string{"image/jpeg", "image/png"})

	router := mux.NewRouter()
	router.Handle("/avatar", authMW.AuthMiddleware(http.HandlerFunc(handler.UploadAvatar))).Methods(http.MethodPost)
	router.Handle("/avatar", authMW.AuthMiddleware(http.HandlerFunc(handler.DeleteAvatar))).Methods(http.MethodDelete)

	userID := uuid.New().String()
	token, _ := jwtManager.Generate(userID)

	t.Run("UploadAvatar - success", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Disposition", `form-data; name="avatar"; filename="test.jpg"`)
		partHeader.Set("Content-Type", "image/jpeg")
		part, _ := writer.CreatePart(partHeader)
		_, _ = io.Copy(part, strings.NewReader("fake-image-data"))
		writer.Close()

		mockUserService.EXPECT().UploadAvatar(gomock.Any(), gomock.Any()).Return(&dto.UploadAvatarResponse{URL: "http://avatar.url"}, nil)

		req := httptest.NewRequest(http.MethodPost, "/avatar", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "http://avatar.url")
	})

	t.Run("UploadAvatar - fail no file", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/avatar", nil)
		req.Header.Set("Content-Type", "multipart/form-data")
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("UploadAvatar - fail service error", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Disposition", `form-data; name="avatar"; filename="test.png"`)
		partHeader.Set("Content-Type", "image/png")
		part, _ := writer.CreatePart(partHeader)
		_, _ = io.Copy(part, strings.NewReader("fake-image-data"))
		writer.Close()

		mockUserService.EXPECT().UploadAvatar(gomock.Any(), gomock.Any()).Return(nil, errors.New("internal server error"))

		req := httptest.NewRequest(http.MethodPost, "/avatar", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("UploadAvatar - fail unauthenticated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/avatar", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("DeleteAvatar - success", func(t *testing.T) {
		mockUserService.EXPECT().DeleteAvatar(gomock.Any(), dto.DeleteAvatarRequest{UserID: userID}).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "deleted")
	})

	t.Run("DeleteAvatar - fail service error", func(t *testing.T) {
		mockUserService.EXPECT().DeleteAvatar(gomock.Any(), gomock.Any()).Return(errors.New("internal server error"))

		req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("DeleteAvatar - fail unauthenticated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_user.NewMockIService(ctrl)
	jwtManager := jwtmanager.NewManager("test-secret", time.Hour)
	authMW := middleware.NewAuthMiddleware(jwtManager)
	handler := NewHandler(mockUserService, jwtManager, nil, nil)

	router := mux.NewRouter()
	router.Handle("/profile", authMW.AuthMiddleware(http.HandlerFunc(handler.UpdateProfile))).Methods(http.MethodPut)

	userID := uuid.New().String()
	token, _ := jwtManager.Generate(userID)

	updateReqPayload := updateProfileRequest{
		Login: "newlogin123",
		Email: "new@email.com",
	}
	body, _ := json.Marshal(updateReqPayload)

	t.Run("success", func(t *testing.T) {
		serviceResponse := &dto.UpdateProfileResponse{
			ID:    userID,
			Login: updateReqPayload.Login,
			Email: updateReqPayload.Email,
		}
		mockUserService.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(serviceResponse, nil)

		req := httptest.NewRequest(http.MethodPut, "/profile", bytes.NewReader(body))
		req.AddCookie(&http.Cookie{Name: sessionTokenCookie, Value: token})
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var respBody dto.UpdateProfileResponse
		err := json.Unmarshal(rr.Body.Bytes(), &respBody)
		require.NoError(t, err)
		assert.Equal(t, userID, respBody.ID)
	})

	t.Run("fail - unauthenticated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/profile", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}
