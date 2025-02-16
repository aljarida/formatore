package dblogic

import (
	"testing"
)

func TestColumnBlueprintsToQueryComp(t *testing.T) {
	cbs := []ColumnBlueprint{{"col_1", "TEXT"}, {"col_2", "TEXT"}, {"col_3", "TEXT"}}
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
	tb := TableBlueprint{"test_table", []ColumnBlueprint{{"col_1", "TEXT"}}}
	expected := joinStrings(
		"CREATE TABLE IF NOT EXISTS test_table (",
		PKeyComp,
		UnixDatetimeComp,
		"col_1 TEXT);")
	result := makeCreateQuery(tb)
	if result != expected {
		t.Fatalf("Results do not match:\n\tExpected %s;\n\tgot %s.", expected, result)
	}
}

func TestCreateTable(t *testing.T) {
	db, err := ConnectToDB()
	if err != nil {
		t.Fatal("Could not connect to database.")
	}
	
	tb := TableBlueprint{"test_table", []ColumnBlueprint{{"col 1", "text"}}}
	err = CreateTable(db, tb)
	if err != nil {
		t.Fatalf("Failed to create table. Error: %v", err)
	}

	err = DropTable(db, tb.Name)
	if err != nil {
		t.Fatal("Failed to drop table.")
	}
}
