package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IstvanN/domafocapp-backend/controllers"
	"github.com/IstvanN/domafocapp-backend/database"
)

var port = ":" + os.Getenv("PORT")

func main() {
	router := controllers.StartupRouter()

	database.StartupPostgres()
	log.Println("domafocapp-backend is up and running on port :8080")
	log.Fatal(http.ListenAndServe(port, router))
}
