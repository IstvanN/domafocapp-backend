package models

import "gorm.io/gorm"

type Fixture struct {
	gorm.Model
	TournamentID  uint    `json:"tournamentID"`
	HomeTeamID    uint    `json:"homeTeamID"`
	AwayTeamID    uint    `json:"awayTeamID"`
	HomeTeamGoals int     `json:"homeTeamGoals"`
	AwayTeamGoals int     `json:"awayTeamGoals"`
	Scores        []Score `json:"scores"`
}
