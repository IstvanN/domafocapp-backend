package repositories

import (
	"log"
	"time"

	"github.com/IstvanN/domafocapp-backend/database"
	"github.com/IstvanN/domafocapp-backend/models"
)

func GetAllTournaments() ([]models.Tournament, error) {
	var tournaments []models.Tournament
	result := database.GetPostgresDB().Find(&tournaments)
	if result.Error != nil {
		return nil, result.Error
	}
	return tournaments, nil
}

func CreateTournament(note string, noOfRounds int) error {
	newTournament := models.Tournament{
		Date:           time.Now(),
		Note:           note,
		NumberOfRounds: noOfRounds,
	}

	result := database.GetPostgresDB().Create(&newTournament)
	if result.Error != nil {
		return result.Error
	}
	log.Println("new tournament added:", newTournament)
	return nil
}
