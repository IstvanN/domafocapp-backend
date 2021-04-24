package models

import "gorm.io/gorm"

type Fixture struct {
	gorm.Model
	TournamentID  uint
	HomeTeamID    uint
	AwayTeamID    uint
	HomeTeamGoals int
	AwayTeamGoals int
	Scores        []Score
}
