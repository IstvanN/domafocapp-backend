package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Date           time.Time
	Note           string
	NumberOfRounds int
	Teams          []Team
}
