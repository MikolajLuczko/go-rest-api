package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// stores pointers to the comments(?) service
type Handler struct {
	Router *mux.Router
}

// return pointer to a new handler
func NewHandler() *Handler {
	return &Handler{}
}

// sets up routes for the app
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "All good")
	})
}
