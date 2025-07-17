package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-http-server/internal/api"
	"simple-http-server/internal/store"
	"simple-http-server/migrations"
)

type Application struct {
	Logger      *log.Logger
	UserHandler *api.UserHandler
	DB          *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// store here

	// handler here
	userHandler := api.NewUserHandler()

	app := &Application{
		Logger:      logger,
		UserHandler: userHandler,
		DB:          pgDB,
	}
	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "status is Available\n")
}
