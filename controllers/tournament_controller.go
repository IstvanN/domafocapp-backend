package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/domafocapp-backend/repositories"
	"github.com/IstvanN/domafocapp-backend/security"
	"github.com/gorilla/mux"
)

func registerTournamentRoutes(router *mux.Router) {
	s := router.PathPrefix("/tournaments").Subrouter()
	s.HandleFunc("", allTournamentsHandler).Methods(http.MethodGet)
}

func allTournamentsHandler(w http.ResponseWriter, r *http.Request) {
	tournaments, err := repositories.GetAllTournaments()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tournaments)
}
