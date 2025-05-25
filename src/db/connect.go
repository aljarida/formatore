package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"formatore/src/enums"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB(dbName string) (*sql.DB, error) {
	execPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Could not determine executable path: %w.", err)
	}

	execDir := filepath.Dir(execPath)
	absDBPath := filepath.Join(execDir, dbName)

	db, err := sql.Open(enums.DriverName, absDBPath)
	if err != nil {
		return nil, fmt.Errorf("Could not create database connection: ~%v~.", err)
	}
	return db, nil
}

