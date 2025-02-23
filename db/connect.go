package db

import (
	"database/sql"
	"fmt"
	"formatore/enums"
	_ "github.com/mattn/go-sqlite3"
)


func ConnectToDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open(enums.DriverName, dbName)
	if err != nil {
		return nil, fmt.Errorf("Could not create database connection: ~%v~.", err)
	}
	return db, nil
}
