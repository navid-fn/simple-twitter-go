package routes

import (
	"simple-http-server/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	// User
	r.Get("/user/{id}", app.UserHandler.HandleGetUserByID)
	r.Post("/user", app.UserHandler.HandleCreateUser)


	// Posts
	r.Get("/post/{id}", app.PostHandler.HandleGetPostByID)
	r.Post("/post", app.PostHandler.HandleCreatePost)

	return r

}
