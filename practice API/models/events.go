package models

import "time"

type Events struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userid"`
	Name        string    `json:"name" binding: "required"`
	Description string    `json:"description" binding: "required"`
	Date        time.Time `json:"date" binding: "required"`
	Location    string    `json:"location" binding: "required"`
}

var events []Events

func GetAllEvents() []Events {
	return events
}