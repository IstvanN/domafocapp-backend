package models

import "gorm.io/gorm"

type Footballer struct {
	gorm.Model
	Name          string
	TeamID        uint
	ParticipantID uint
}
