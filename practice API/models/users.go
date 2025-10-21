package models

import (
	"RESTApi/db"
	"RESTApi/utils"
	"errors"
)

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding: "required"`
	Password string `json:"password" binding: "required"`
}

func (u *Users) Save() error {
	insertQuery:= `INSERT INTO users (email, password) VALUES (?,?)`
	HashedPassword, err:= utils.HashPassword(u.Password)
	if err != nil{
		return err
	}

	_, err = db.DB.Exec(insertQuery, u.Email, HashedPassword)
	if err != nil {
		return errors.New("unable to create table")
	}

	return nil

}


