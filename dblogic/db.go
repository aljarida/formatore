package dblogic

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const DBName = "formatore.db"
const DriverName = "sqlite3"

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open(DriverName, DBName)
	if err != nil {
		return nil, errors.New("Could not create database connection.")
	}
	return db, nil
}
