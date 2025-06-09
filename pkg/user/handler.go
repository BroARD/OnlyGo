package user

import (
	"OnlyGo/pkg/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service UserService
}

func NewHandler(serv UserService) handlers.Handler {
	return &userHandler{service: serv}
}

func (h *userHandler) Register(router *mux.Router) {
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users", h.GetUsers).Methods("GET")
}
 
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user NewUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	err = h.service.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.service.GetUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Could not transform from json to bytes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

