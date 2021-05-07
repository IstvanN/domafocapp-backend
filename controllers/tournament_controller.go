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
	s.HandleFunc("/create", createTournament).Methods(http.MethodPost)
}

func allTournamentsHandler(w http.ResponseWriter, r *http.Request) {
	tournaments, err := repositories.GetAllTournaments()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tournaments)
}

func createTournament(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Note           string `json:"note"`
		NumberOfRounds int    `json:"numberOfRounds"`
		Teams          []struct {
			TeamName    string `json:"teamName"`
			Footballers []struct {
				FootballerName string `json:"footballerName"`
				ParticipantID  uint   `json:"participantID"`
			} `json:"footballers"`
		} `json:"teams"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.NumberOfRounds == 0 {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.CreateTournament(rb.Note, rb.NumberOfRounds, rb.Teams); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "tournament created successfully")
}
