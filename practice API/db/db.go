package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}

func CreateTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS events  (
     id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL ,
    description TEXT NOT NULL ,
    location    TEXT NOT NULL ,
    dateTime    DATETIME NOT NULL ,
    userId INTEGER NOT NULL ,
    
     )`

	 _, err:= DB.Exec(createTableQuery) 

	 if err != nil{
		panic("Not able to create events.table")
	 } 


}