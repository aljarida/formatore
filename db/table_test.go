package db

import (
	"database/sql"
	"formatore/enums"
	"formatore/structs"
	"formatore/utils"
	"reflect"
	"testing"
)

func TestColumnBlueprintsToQueryComp(t *testing.T) {
	cbs := []structs.ColumnBlueprint{
		{Name: "col_1", Type: enums.Text},
		{Name: "col_2", Type: enums.Text},
		{Name: "col_3", Type: enums.Text},
	}
	expected := "col_1 TEXT, col_2 TEXT, col_3 TEXT"
	result := columnBlueprintsToQueryComp(cbs)
	if result != expected {
		t.Fatalf("Results do not match: Expected %s; got %s.", expected, result)
	}

	expected = ""
	result = columnBlueprintsToQueryComp(nil)
	if result != expected {
		t.Fatalf("Results do not match: Expected %s; got %s.", expected, result)
	}
}

func TestMakeCreateQuery(t *testing.T) {
	tb := structs.TableBlueprint{
		Name:             "test_table",
		ColumnBlueprints: []structs.ColumnBlueprint{{Name: "col_1", Type: enums.Text}},
	}
	expected := utils.JoinStrings(
		"CREATE TABLE IF NOT EXISTS test_table (",
		enums.PKeyComp,
		enums.UnixDatetimeComp,
		"col_1 TEXT);")
	result := makeCreateQuery(tb)
	if result != expected {
		t.Fatalf("Results do not match:\n\tExpected %s;\n\tgot %s.", expected, result)
	}
}

func TestCreateTable(t *testing.T) {
	db := connectToTestDB(t)
	tb := structs.TableBlueprint{
		Name:             "test_table",
		ColumnBlueprints: []structs.ColumnBlueprint{{Name: "col_1", Type: "text"}},
	}
	createTestTable(t, db, tb)
	dropTestTable(t, db, tb.Name)
}

func TestDropTable(t *testing.T) {
	db := connectToTestDB(t)

	tableName := "table_to_drop"
	tb := structs.TableBlueprint{
		Name:             tableName,
		ColumnBlueprints: []structs.ColumnBlueprint{{Name: "col 1", Type: "text"}},
	}

	createTestTable(t, db, tb)

	tableNames, err := TableNames(db)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.Has(tableNames, tableName) {
		t.Fatalf("Expected database to contain table '%s'but it did not.", tableName)
	}

	dropTestTable(t, db, tb.Name)

	tableNames, err = TableNames(db)
	if err != nil {
		t.Fatal(err)
	}

	if utils.Has(tableNames, tableName) {
		t.Fatalf("Expected database not to contain table '%s' but it did.", tableName)
	}
}

// Helper function for TestDropAllTables and TestTableNames.
func createManyTables(t *testing.T, db *sql.DB, names []string) {
	cbs := []structs.ColumnBlueprint{{Name: "col_1", Type: enums.Text}}
	for _, n := range names {
		tb := structs.TableBlueprint{
			Name:             n,
			ColumnBlueprints: cbs,
		}
		err := CreateTable(db, tb)
		if err != nil {
			t.Fatalf("Failed to create table. Error: %v", err)
		}
	}
}

func TestDropAllTables(t *testing.T) {
	db := connectToTestDB(t)

	names := []string{"table_one", "table_two", "table_three"}
	createManyTables(t, db, names)

	err := DropAllTables(db)
	if err != nil {
		t.Fatalf("Error dropping tables: ~%v~.", err)
	}

	existingTables, err := TableNames(db)
	if err != nil {
		t.Fatalf("Error obtaining table names: ~%v~.", err)
	}

	if len(existingTables) > 0 {
		t.Fatalf("Failed to drop tables: `%v`.", existingTables)
	}
}

func TestTableNames(t *testing.T) {
	db := connectToTestDB(t)

	names := []string{"table_one", "table_two", "table_three"}
	createManyTables(t, db, names)
	existingTables, err := TableNames(db)
	if err != nil {
		t.Fatalf("Error obtaining table names: ~%v~.", err)
	}

	if !reflect.DeepEqual(names, existingTables) {
		t.Fatalf("Expected tables added (%v) to match returned table names (%v).", names, existingTables)
	}

	err = DropAllTables(db)
	if err != nil {
		t.Fatalf("Error dropping tables: ~%v~.", err)
	}
}
