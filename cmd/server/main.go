package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/MikolajLuczko/go-rest-api/internal/transport/http"
)

// the App struct will contain things like pointers to db connections CHANGE THIS COMMENT LATER
type App struct{}

// sets up the application
func (app *App) Run() error {
	fmt.Println("Setting up the API")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
	}
	return nil
}

func main() {
	fmt.Println("Welcome to this GO REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the REST API")
		fmt.Println(err)
	}
}
