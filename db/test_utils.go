package db

import (
	"database/sql"
	"formatore/enums"
	"formatore/structs"
	"testing"
)

func connectToTestDB(t *testing.T) *sql.DB {
	db, err := ConnectToDB(enums.TestDBName)
	if err != nil {
		t.Fatalf("Could not connect to database: ~%v~", err)
	}

	return db
}

func createTestTable(t *testing.T, db *sql.DB, tb structs.TableBlueprint) {
	err := CreateTable(db, tb)
	if err != nil {
		t.Fatalf("Failed to create table. Error: ~%v~.", err)
	}
}

func dropTestTable(t *testing.T, db *sql.DB, name string) {
	err := DropTable(db, name)
	if err != nil {
		t.Fatalf("Failed to drop table. Error: ~%v~.", err)
	}
}
