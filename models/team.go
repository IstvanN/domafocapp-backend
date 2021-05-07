package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name         string       `json:"name"`
	TournamentID uint         `json:"tournamentID"`
	Footballers  []Footballer `json:"footballers"`
}
