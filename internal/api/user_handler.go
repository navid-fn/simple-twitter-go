package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
	fmt.Fprintf(w, "User Created.")
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
