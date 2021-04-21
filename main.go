package main

import (
	"log"
	"net/http"

	"github.com/IstvanN/domafocapp-backend/database"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to DomafocApp!"))
	})

	database.StartupPostgres()
	log.Println("domafocapp-backend is up and running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
