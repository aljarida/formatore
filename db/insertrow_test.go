package db

import (
	"database/sql"
	"formatore/enums"
	"formatore/structs"
	"reflect"
	"testing"
)

const testTableName = "MY_TABLE"

var expectedCbs = []structs.ColumnBlueprint{
	{Name: enums.PKeyFieldName, Type: enums.Integer},
	{Name: enums.UnixDataTimeFieldName, Type: enums.Integer},
	{Name: "name", Type: enums.Text},
	{Name: "weight", Type: enums.Real},
	{Name: "height", Type: enums.Integer},
}

func setupTable(t *testing.T, db *sql.DB) {
	cbs := []structs.ColumnBlueprint{
		{Name: "name", Type: enums.Text},
		{Name: "weight", Type: enums.Real},
		{Name: "height", Type: enums.Integer},
	}

	tb := structs.TableBlueprint{
		Name:             testTableName,
		ColumnBlueprints: cbs,
	}

	createTestTable(t, db, tb)
}

func TestExtractColumnNames(t *testing.T) {
	expected := []string{
		enums.PKeyFieldName,
		enums.UnixDataTimeFieldName,
		"name",
		"weight",
		"height",
	}
	result := extractColumnNames(expectedCbs)
	if !reflect.DeepEqual(result, expected) {
		t.Logf("Lengths of lists are %d and %d:", len(expected), len(result))
		t.Fatalf("Expected %v but got %v.", expected, result)
	}

	cbs := []structs.ColumnBlueprint{}
	expected = []string{}
	result = extractColumnNames(cbs)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v.", expected, result)
	}
}

func TestColumnBlueprints(t *testing.T) {
	db := connectToTestDB(t)
	setupTable(t, db)
	defer dropTestTable(t, db, testTableName)

	result, err := ColumnBlueprints(db, testTableName)
	if err != nil {
		t.Fatalf("Error obtaining columns: ~%v~.", err)
	}

	if !reflect.DeepEqual(expectedCbs, result) {
		t.Fatalf("Expected `%v` but got `%v`.", expectedCbs, result)
	}
}

func TestInsertRow(t *testing.T) {
	db := connectToTestDB(t)
	setupTable(t, db)
	defer dropTestTable(t, db, testTableName)

	values := []string{"$", "1.0", "1"}
	err := InsertRow(db, testTableName, values)
	if err != nil {
		t.Fatalf("Error w/ InsertRow: ~%v~.", err)
	}
}
