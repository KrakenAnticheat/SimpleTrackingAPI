package repository

import (
	"simpletrackingapi/internal/database"
	"simpletrackingapi/internal/models"
)

func SaveEvent(event models.Event) {
	database.Events = append(database.Events, event)
}

func GetAllEvents() []models.Event {
	return database.Events
}