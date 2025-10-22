package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"spotify/internal/middleware"
	"spotify/internal/user/dto"
	"spotify/pkg/response"
	"strings"
	"time"
)

const sessionTokenCookie = "session_token"

type registerRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *registerRequest) validate() error {
	if len(i.Login) < 5 {
		return fmt.Errorf("login is too short (minimum 5 chars)")
	}
	if len(i.Password) < 8 {
		return fmt.Errorf("password is too short (minimum 8 chars)")
	}
	if !strings.Contains(i.Email, "@") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

type registerResponse struct {
	ID string `json:"id"`
}
type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i *loginRequest) validate() error {
	if i.Login == "" {
		return fmt.Errorf("empty login")
	}
	if i.Password == "" {
		return fmt.Errorf("empty password")
	}
	return nil
}

type loginResponse struct {
	ID string `json:"id"`
}

type logoutResponse struct {
	Status string `json:"status"`
}

type uploadAvatarRequest struct {
	ContentType string
	Size        int64
}

func (r *uploadAvatarRequest) validate() error {
	if r.Size == 0 {
		return fmt.Errorf("empty file")
	}
	if r.Size > 5<<20 {
		return fmt.Errorf("max 5MB")
	}
	switch r.ContentType {
	case "image/png", "image/jpeg":
		return nil
	default:
		return fmt.Errorf("unsupported content type: %s", r.ContentType)
	}
}

type uploadAvatarResponse struct {
	URL string `json:"avatar_url"`
}
type deleteAvatarResponse struct {
	Status string `json:"status"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	const op = "[Register] "
	defer r.Body.Close()

	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("%s invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	if err := req.validate(); err != nil {
		log.Printf("%s validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	user, err := h.svc.Register(r.Context(), dto.RegisterRequest{
		Login:    req.Login,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("%s service error: %v", op, err)
		handleServiceError(w, err)
		return
	}

	token, err := h.jwtManager.Generate(user.ID)
	if err != nil {
		log.Printf("%s ERROR: failed to generate token", op)
		response.InternalErrorJSON(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    token,
		Expires:  time.Now().Add(h.jwtManager.GetTTL()),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusCreated, registerResponse{ID: user.ID})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	const op = "[Login] "
	defer r.Body.Close()

	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("%s invalid body: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	if err := req.validate(); err != nil {
		log.Printf("%s validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	user, err := h.svc.Login(r.Context(), dto.LoginRequest{
		Login:    req.Login,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("%s service error:: %v", op, err)
		handleServiceError(w, err)
		return
	}

	token, err := h.jwtManager.Generate(user.ID)
	if err != nil {
		log.Printf("%s ERROR: failed to generate token", op)
		response.InternalErrorJSON(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    token,
		Expires:  time.Now().Add(h.jwtManager.GetTTL()),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, loginResponse{ID: user.ID})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookie,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	response.JSON(w, http.StatusOK, logoutResponse{Status: "ok"})
}

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	const op = "[UploadAvatar] "

	userID, ok := middleware.GetUserID(r.Context())
	if !ok || userID == "" {
		response.UnauthorizedJSON(w)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		log.Printf("%s failed to get file: %v", op, err)
		response.BadRequestJSON(w)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("%s failed to read file: %v", op, err)
		response.InternalErrorJSON(w)
		return
	}

	req := uploadAvatarRequest{
		ContentType: header.Header.Get("Content-Type"),
		Size:        header.Size,
	}

	if err := req.validate(); err != nil {
		log.Printf("%s validation error: %v", op, err)
		response.BadRequestJSON(w)
		return
	}

	res, err := h.svc.UploadAvatar(r.Context(), dto.UploadAvatarRequest{
		UserID:      userID,
		File:        fileBytes,
		ContentType: req.ContentType,
	})
	if err != nil {
		log.Printf("%s service error: %v", op, err)
		handleServiceError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, uploadAvatarResponse{URL: res.URL})
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	const op = "[DeleteAvatar] "

	userID, ok := middleware.GetUserID(r.Context())
	if !ok || userID == "" {
		response.UnauthorizedJSON(w)
		return
	}

	req := dto.DeleteAvatarRequest{
		UserID: userID,
	}

	if err := h.svc.DeleteAvatar(r.Context(), req); err != nil {
		log.Printf("%s service error: %v", op, err)
		handleServiceError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, deleteAvatarResponse{Status: "deleted"})
}
