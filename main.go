package main

import (
	"flag"
	"fmt"
	"net/http"
	"simple-http-server/internal/app"
	"simple-http-server/internal/routes"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "backend port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	app.Logger.Printf("Running App. Port: %d", port)
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
