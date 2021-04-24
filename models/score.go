package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	FootballerID uint
	FixtureID    uint
}
