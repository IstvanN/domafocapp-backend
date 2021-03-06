package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler).Methods(http.MethodGet, http.MethodOptions)
	registerTournamentRoutes(router)
	return router
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to DomafocApp!"))
}

func writeMessage(w http.ResponseWriter, msg string) {
	finalMessage := fmt.Sprintf("{\"message\": \"%s\"}", msg)
	w.Write([]byte(finalMessage))
}
