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
	PostHandler *api.PostHandler
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
	postStore := store.NewPostgresPostStore(pgDB)
	userStore := store.NewPostgresUserStore(pgDB)

	// handler here
	userHandler := api.NewUserHandler(userStore)
	postHandler := api.NewPostHandler(postStore)

	app := &Application{
		Logger:      logger,
		UserHandler: userHandler,
		PostHandler: postHandler,
		DB:          pgDB,
	}
	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "status is Available\n")
}
