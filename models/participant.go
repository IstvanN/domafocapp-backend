package models

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	Nickname string `json:"nickname"`
}
