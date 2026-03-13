// DEMO DB
package database

import "simpletrackingapi/internal/models"

var Events []models.Event

func InitDB() {
	Events = []models.Event{}
}