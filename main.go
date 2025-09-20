package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Handlers struct {
	users []User
	mu    *sync.Mutex
}

func NewHandler() *Handlers {
	return &Handlers{
		users: make([]User, 0),
		mu:    &sync.Mutex{},
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserInput := new(UserInput)

	err := decoder.Decode(newUserInput)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	if len(newUserInput.Name) < 5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "name is too short (minimum 5 chars)"})
		return
	}

	if len(newUserInput.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "password is too short (minimum 8 chars)"})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, user := range h.users {
		if user.Name == newUserInput.Name {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "username alredy exists"})
			return
		}

		if user.Email == newUserInput.Email {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "email alredy used"})
			return
		}
	}

	var id uint64 = 0
	if len(h.users) > 0 {
		id = h.users[len(h.users)-1].ID + 1
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserInput.Password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "server error"})
		return
	}

	h.users = append(h.users, User{
		ID:       id,
		Name:     newUserInput.Name,
		Email:    newUserInput.Email,
		Password: string(hashedPassword),
	})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]uint64{"id": id})
}

func (h *Handlers) AutorizationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newloginUser := new(LoginInput)

	err := decoder.Decode(newloginUser)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	var foundUser *User
	for i := range h.users {
		if h.users[i].Name == newloginUser.Name {
			foundUser = &h.users[i]
			break
		}
	}

	if foundUser == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "incorrect username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(newloginUser.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "incorrect username or password"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]uint64{"id": foundUser.ID})

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "This is the Spotify home page"}`))
}

func newRouter(h *Handlers) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/registration/", h.RegisterHandler).Methods("POST")
	r.HandleFunc("/autorization/", h.AutorizationHandler).Methods("POST")
	return r
}
func runServer(r *mux.Router, addr string) {
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	fmt.Println("starting server at", addr)
	fmt.Println("  GET  http://localhost:8080/")
	server.ListenAndServe()
}

func main() {
	handler := NewHandler()
	router := newRouter(handler)

	runServer(router, ":8080")
}
