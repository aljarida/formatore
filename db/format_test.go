package db

import (
	"formatore/structs"
	"reflect"
	"testing"
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
	cbs := []structs.ColumnBlueprint{
		{Name: "col 1", Type: "text"},
		{Name: "col 2", Type: "integer"},
	}
	expected := []structs.ColumnBlueprint{
		{Name: "col_1", Type: "TEXT"},
		{Name: "col_2", Type: "INTEGER"},
	}
	formatColumnBlueprints(cbs)
	result := cbs
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Result was not formatted correctly.")
	}
}
