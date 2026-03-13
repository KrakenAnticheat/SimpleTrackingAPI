package handlers

import (
	"encoding/json"
	"net/http"

	"simpletrackingapi/internal/models"
	"simpletrackingapi/internal/services"
	"simpletrackingapi/pkg/utils"
)

func TrackEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.JSON(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, "invalid json")
		return
	}

	err = services.ProcessEvent(event)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, "event tracked")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {

	events := services.GetEvents()

	utils.JSON(w, http.StatusOK, events)

}