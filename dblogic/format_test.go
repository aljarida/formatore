package dblogic

import (
	"testing"
	"reflect"
)

func TestFormatCBName(t *testing.T) {
	cbName := "FIRST COLUMN"
	expected := "first_column"
	result := formatCBName(cbName)
	if result != expected {
		t.Fatal("Results do not match.")
	}
}

func TestFormatCBType(t *testing.T) {
	cbType := "integer"
	expected := "INTEGER"
	result := formatCBType(cbType)
	if result != expected {
		t.Fatal("Results do not match.")
	}
}

func TestFormatColumnBlueprints(t *testing.T) {
	cbs := []ColumnBlueprint{{"col 1", "text"}, {"col 2", "integer"}}
	result := formatColumnBlueprints(cbs)
	expected := []ColumnBlueprint{{"col_1", "TEXT"}, {"col_2", "INTEGER"}}
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Result was not formatted correctly.")
	}
}
