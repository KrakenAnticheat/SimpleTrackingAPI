package models

type Event struct {
	UserID    string                 `json:"user_id"`
	EventName string                 `json:"event"`
	Timestamp int64                  `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata"`
}