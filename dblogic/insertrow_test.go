package dblogic

import (
	"testing"
	"reflect"
	"database/sql"
	"formatore/utils"
)

const testTableName = "MY_TABLE"
var metadata = []utils.ColumnBlueprint{
	{utils.PKeyFieldName, utils.Integer},
	{utils.UnixDataTimeFieldName, utils.Integer},
	{"name", utils.Text},
	{"weight", utils.Real},
	{"height", utils.Integer},
}

func setupTable(t *testing.T, db *sql.DB) {
	cbs := []utils.ColumnBlueprint{{"name", utils.Text}, {"weight", utils.Real}, {"height", utils.Integer}}
	tb := utils.TableBlueprint{testTableName, cbs}
	if err := CreateTable(db, tb); err != nil {
		t.Fatal(err)
	}
}

func teardownTable(t *testing.T, db  *sql.DB) {
	if err := DropTable(db, testTableName); err != nil {
		t.Fatal(err)
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
	cbs := []utils.ColumnBlueprint{{"name", utils.Text}, {"weight", utils.Real}, {"height", utils.Integer}}
	expected := []string{"name", "weight", "height"}
	result := getColumnNames(cbs)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v.", expected, result)
	}

	cbs = []utils.ColumnBlueprint{}
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
