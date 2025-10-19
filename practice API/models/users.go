package models

import (
	"RESTApi/db"
	"errors"
)

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding: "required"`
	Password string `json:"email" binding: "required"`
}

func (u *Users) Save() error {
	insertQuery:= `INSERT INTO users (email, password) VALUES (?,?)`
	_, err := db.DB.Exec(insertQuery, u.Email, u.Password)
	if err != nil {
		return errors.New("unable to create table")
	}

	return nil
}
