package router

import (
	apphandlers "github.com/Oleg-Skalozub/testtask/src/app-handlers"
	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	appHandlers := apphandlers.NewHandler()
	router.HandleFunc("/request", appHandlers.Request).Methods("GET")

	return router
}
