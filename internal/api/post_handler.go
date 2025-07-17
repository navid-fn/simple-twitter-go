package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func (uh *PostHandler) HandleGetPostByID(w http.ResponseWriter, r *http.Request) {
	paramPostID := chi.URLParam(r, "id")
	if paramPostID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramPostID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Post ID data: %d", userID)
}

func (uh *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Created.")
}

func (uh *PostHandler) HandleUpdatePostByID(w http.ResponseWriter, r *http.Request) {
	paramPostID := chi.URLParam(r, "id")
	if paramPostID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramPostID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Post ID data: %d", userID)

}

func (uh *PostHandler) HandleDeletePostByID(w http.ResponseWriter, r *http.Request) {
paramPostID := chi.URLParam(r, "id")
	if paramPostID == "" {
		http.NotFound(w, r)
		return
	}
	userID, err := strconv.ParseInt(paramPostID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Post ID data: %d", userID)
}
