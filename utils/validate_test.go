package utils

import (
	"formatore/enums"
	"formatore/structs"
	"testing"
)

var metadata = []structs.ColumnBlueprint{
	{Name: enums.PKeyFieldName, Type: enums.Integer},
	{Name: enums.UnixDataTimeFieldName, Type: enums.Integer},
	{Name: "name", Type: enums.Text},
	{Name: "weight", Type: enums.Real},
	{Name: "height", Type: enums.Integer},
}

func TestValidateTableBlueprint(t *testing.T) {
	var err error
	tbValid := structs.TableBlueprint{
		Name:             "test_table",
		ColumnBlueprints: []structs.ColumnBlueprint{{Name: "col_1", Type: enums.Text}},
	}
	err = ValidateTableBlueprint(tbValid)
	if err != nil {
		t.Fatal("Error thrown for valid TableBlueprint.")
	}

	tbNoColumnBlueprints := structs.TableBlueprint{Name: "test_table", ColumnBlueprints: nil}
	err = ValidateTableBlueprint(tbNoColumnBlueprints)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}

	tbNoName := structs.TableBlueprint{
		Name:             "",
		ColumnBlueprints: []structs.ColumnBlueprint{{Name: "col 1", Type: "TEXT"}},
	}
	err = ValidateTableBlueprint(tbNoName)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}
}

func TestValidateAndApostrophizeValues(t *testing.T) {
	var err error
	values := []string{"1", "$", "1.0", "1"}
	if err = ValidateAndApostrophizeValues(metadata, values); err != nil {
		t.Fatalf("Encountered error with valid values %v: ~%v~.", values, err)
	}

	if values[1] != "'$'" {
		t.Fatalf("enums.Text was not apostrophized.")
	}

	values = []string{"1", "1", "1.0", "1"}
	if err = ValidateAndApostrophizeValues(metadata, values); err == nil {
		t.Fatalf("Executed without error despite invalid values %v: ~%v~.", values, err)
	}

	values = []string{"1", "$", "1.0"}
	if err = ValidateAndApostrophizeValues(metadata, values); err == nil {
		t.Fatalf("Executed without error despite too few values %v: ~%v~.", values, err)
	}
}
