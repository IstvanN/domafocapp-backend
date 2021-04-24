package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IstvanN/domafocapp-backend/models"
	"github.com/IstvanN/domafocapp-backend/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresDB  *gorm.DB
	postgresURL = os.Getenv("DATABASE_URL")
)

// StartupPostgres is the init call of the Postgres DB, supposed to be called in the main function
func StartupPostgres() {
	if postgresURL == "" {
		log.Fatalln("unable to connect to Postgres DB: no URL provided")
		return
	}
	dbHost, dbUser, dbPort, dbPw, dbName := services.GetPostgresDBSpecsFromURL(postgresURL)
	dbSpecs := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v TimeZone=Europe/Budapest",
		dbHost, dbPort, dbUser, dbPw, dbName)

	var err error
	postgresDB, err = gorm.Open(postgres.Open(dbSpecs), &gorm.Config{})
	if err != nil {
		try := 1
		for try <= 6 && err != nil {
			log.Printf("establishing connection to the database... %d\nExiting after 5 tries.", try)
			time.Sleep(2 * time.Second)
			postgresDB, err = gorm.Open(postgres.Open(dbSpecs), &gorm.Config{})
			try++
			if try == 6 {
				panic(err)
			}
		}
	}

	postgresDB.AutoMigrate(&models.Tournament{})
	postgresDB.AutoMigrate(&models.Team{})

	log.Println("successfully connected to Postgres DB!")
}

// GetPostgresDB is the conventional function to access the DB session
func GetPostgresDB() *gorm.DB {
	return postgresDB
}
