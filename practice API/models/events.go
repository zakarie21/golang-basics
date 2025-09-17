package models

import (
	//"fmt"
	"RESTApi/db"
	"errors"
	"time"
)

type Events struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userid"`
	Name        string    `json:"name" binding: "required"`
	Description string    `json:"description" binding: "required"`
	Date        time.Time `json:"date" binding: "required"`
	Location    string    `json:"location" binding: "required"`
}

var events []Events

func GetAllEvents() ([]Events, error) {
	rows, err:= db.DB.Query(`SELECT * FROM events`)
	if err != nil{
		return events, errors.New("command failed to exectue")
	}
	
	for rows.Next(){
		var event Events
		err:= rows.Scan(&event.ID,&event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return events, errors.New("unable to scan rows")
		}
		events = append(events, event)
	}

	return events, nil
}

func (e  *Events) Save() error {
	insertQuery:= `INSERT INTO events (name, description,location,dateTime, userId) VALUES (?,?,?,?,?)`
	_, err:= db.DB.Exec(insertQuery, e.Name, e.Description, e.Location, e.Date,e.UserID)
	if err!= nil{
		return errors.New("unable to create table")
	}

	return nil
}

