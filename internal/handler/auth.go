package handler

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"spotify/internal/model"
	"spotify/pkg/response"
)

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newUserInput := new(model.UserInput)

	if err := json.NewDecoder(r.Body).Decode(&newUserInput); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "invalid request body"})
		return
	}

	if err := newUserInput.ValidateUserInput(); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, user := range h.users {
		if user.Name == newUserInput.Name {
			response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "username alredy exists"})
			return
		}

		if user.Email == newUserInput.Email {
			response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "email alredy used"})
			return
		}
	}

	var id uint64 = 0
	if len(h.users) > 0 {
		id = h.users[len(h.users)-1].ID + 1
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserInput.Password), bcrypt.DefaultCost)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: "server error"})
		return
	}

	h.users = append(h.users, model.User{
		ID:       id,
		Name:     newUserInput.Name,
		Email:    newUserInput.Email,
		Password: string(hashedPassword),
	})
	response.JSON(w, http.StatusCreated, map[string]uint64{"id": id})
}

func (h *Handlers) AutorizationHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newloginUser := new(model.LoginInput)

	if err := json.NewDecoder(r.Body).Decode(&newloginUser); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "invalid request body"})
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	var foundUser *model.User
	for i := range h.users {
		if h.users[i].Name == newloginUser.Name {
			foundUser = &h.users[i]
			break
		}
	}

	if foundUser == nil {
		response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(newloginUser.Password)); err != nil {
		response.JSON(w, http.StatusUnauthorized, response.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	response.JSON(w, http.StatusOK, map[string]uint64{"id": foundUser.ID})

}
