package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-http-server/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewUserHandler(userStore store.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (uh *UserHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	paramUserID := chi.URLParam(r, "id")
	if paramUserID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramUserID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "User ID data: %d", userID)
}

func (uh *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error Has occurred...", err)
		http.Error(w, "Error has occurred..", http.StatusInternalServerError)
	}
	
	createdUser, err := uh.userStore.CreateUser(&user)
	if err != nil {
		fmt.Println("Error Has occurred...", err)
		http.Error(w, "Error has occurred..", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
	return
	
}

func (uh *UserHandler) HandleUpdateUserByID(w http.ResponseWriter, r *http.Request) {
	paramUserID := chi.URLParam(r, "id")
	if paramUserID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramUserID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "User ID data: %d", userID)

}

func (uh *UserHandler) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
paramUserID := chi.URLParam(r, "id")
	if paramUserID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramUserID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "User ID data: %d", userID)
}
