package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-http-server/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postStore store.PostStore
}

func NewPostHandler(postStore store.PostStore) *PostHandler {
	return &PostHandler{
		postStore: postStore,
	}
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

func (ph *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post store.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		http.Error(w, "Error has been Occurred...", http.StatusInternalServerError)
	}
	createdPost, err := ph.postStore.CreatePost(&post)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		http.Error(w, "Error has been Occurred...", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPost)
	return

}

func (ph *PostHandler) HandleUpdatePostByID(w http.ResponseWriter, r *http.Request) {
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

func (ph *PostHandler) HandleDeletePostByID(w http.ResponseWriter, r *http.Request) {
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
