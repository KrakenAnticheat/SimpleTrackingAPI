package services

import (
	"errors"
	"time"

	"simpletrackingapi/internal/models"
	"simpletrackingapi/internal/handler/repository"
)

func ProcessEvent(event models.Event) error {

	if event.UserID == "" || event.EventName == "" {
		return errors.New("invalid event")
	}

	if event.Timestamp == 0 {
		event.Timestamp = time.Now().Unix()
	}

	repository.SaveEvent(event)

	return nil
}

func GetEvents() []models.Event {
	return repository.GetAllEvents()
}