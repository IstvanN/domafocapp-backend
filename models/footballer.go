package models

import "gorm.io/gorm"

type Footballer struct {
	gorm.Model
	Name          string `json:"name"`
	TeamID        uint   `json:"teamID"`
	ParticipantID uint   `json:"participantID"`
}
