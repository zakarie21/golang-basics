package db

import (
	_"github.com/mattn/go-sqlite3"
	"database/sql"
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
}