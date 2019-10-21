package router

import (
	"github.com/Oleg-Skalozub/testtask/src/app-handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	appHandlers := apphandlers.NewHandler()
	router.HandleFunc("/request", appHandlers.Request).Methods(http.MethodGet)

	return router
}
