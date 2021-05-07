package models

import (
	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Date           string    `json:"date"`
	Note           string    `json:"note"`
	NumberOfRounds int       `json:"numberOfRounds"`
	Teams          []Team    `json:"teams"`
	Fixtures       []Fixture `json:"fixtures"`
}
