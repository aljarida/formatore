package dblogic

import (
	"testing"
	"reflect"
	"database/sql"
)

const testTableName = "MY_TABLE"
var metadata = []ColumnBlueprint{
	{PKeyFieldName, Integer},
	{UnixDataTimeFieldName, Integer},
	{"name", Text},
	{"weight", Real},
	{"height", Integer},
}

func setupTable(t *testing.T, db *sql.DB) {
	cbs := []ColumnBlueprint{{"name", Text}, {"weight", Real}, {"height", Integer}}
	tb := TableBlueprint{testTableName, cbs}
	if err := CreateTable(db, tb); err != nil {
		t.Fatal(err)
	}
}

func teardownTable(t *testing.T, db  *sql.DB) {
	if err := DropTable(db, testTableName); err != nil {
		t.Fatal(err)
	}
}

func TestInferType(t *testing.T) {
	values := []string{
		"1", "-1",
		"1.", "-1.",
		" 1 ", " - 1 ",
		"1.1", "-1.1",
		".1", "-.1",
		" . 1 ", " - . 1 ",
		"--1", "+ 1.1",
		"+1", "+1.1",
		"11111.1", "1.11111",
		"1 . 1", "1 . 1.",
		".", "$",
		"0", "0,1",
		"00", "0.0",
	}

	expected := []string{
		Integer, Integer,
		Real, Real,
		Text, Text,
		Real, Real,
		Real, Real,
		Text, Text,	
		Text, Text,
		Integer, Real, 
		Real, Real,
		Text, Text,
		Text, Text,
		Integer, Text,
		Integer, Real,
	}

	for i, value := range values {
		if res := InferType(value); res != expected[i] {
			t.Fatalf("Expected %s but inferred %s from '%s'.", expected[i], res, values[i])
		}
	}
}

func TestValidateAndApostrophizeValues(t *testing.T) {
	var err error
	values := []string{"1", "$", "1.0", "1"}
	if err = validateAndApostrophizeValues(metadata, values); err != nil {
		t.Fatalf("Encountered error with valid values %v: ~%v~.", values, err)
	}

	if values[1] != "'$'" {
		t.Fatalf("Text was not apostrophized.")
	}
	
	values = []string{"1", "1", "1.0", "1"}
	if err = validateAndApostrophizeValues(metadata, values); err == nil {
		t.Fatalf("Executed without error despite invalid values %v: ~%v~.", values, err)
	}

	values = []string{"1", "$", "1.0"}
	if err = validateAndApostrophizeValues(metadata, values); err == nil {
		t.Fatalf("Executed without error despite too few values %v: ~%v~.", values, err)
	}
}

func TestGetColumnNames(t *testing.T) {
	cbs := []ColumnBlueprint{{"name", Text}, {"weight", Real}, {"height", Integer}}
	expected := []string{"name", "weight", "height"}
	result := getColumnNames(cbs)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v.", expected, result)
	}

	cbs = []ColumnBlueprint{}
	expected = []string{}
	result = getColumnNames(cbs)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v.", expected, result)
	}
}

func TestGetColumnMetadata(t *testing.T) {
	db, err := ConnectToDB()
	if err != nil {
		t.Fatal(err)
	}

	setupTable(t, db)
	defer teardownTable(t, db)


	result, err := getColumnMetadata(db, testTableName)
	if err != nil {
		t.Fatalf("Error obtaining columns: ~%v~.", err)
	}
	
	if !reflect.DeepEqual(result, metadata) {
		t.Fatalf("Result did not match expected: %s vs. %s.", result, metadata)
	}
}

func TestInsertRow(t *testing.T) {
	db, err := ConnectToDB()
	if err != nil {
		t.Fatal(err)
	}

	setupTable(t, db)
	defer teardownTable(t, db)

	values := []string{"1", "$", "1.0", "1"}
	err = InsertRow(db, testTableName, values)
	if err != nil {
		t.Fatalf("Error w/ InsertRow: ~%v~.", err)
	}
}
