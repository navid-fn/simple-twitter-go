package routes

import (
	"simple-http-server/internal/api"
	"simple-http-server/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	// User
	r.Get("/user/{id}", api.NewUserHandler().HandleGetUserByID)
	r.Post("/user/create", api.NewUserHandler().HandleCreateUser)

	return r

}
