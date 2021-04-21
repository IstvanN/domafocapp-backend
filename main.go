package main

import (
	"log"
	"net/http"

	"github.com/IstvanN/domafocapp-backend/controllers"
	"github.com/IstvanN/domafocapp-backend/database"
)

func main() {
	router := controllers.StartupRouter()

	database.StartupPostgres()
	log.Println("domafocapp-backend is up and running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
