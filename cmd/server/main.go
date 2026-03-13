package main

import (
	"log"
	"net/http"

	"simpletrackingapi/internal/database"
	"simpletrackingapi/internal/handlers"
)

func main() {

	database.InitDB()

	http.HandleFunc("/track", handlers.TrackEvent)
	http.HandleFunc("/events", handlers.GetEvents)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}