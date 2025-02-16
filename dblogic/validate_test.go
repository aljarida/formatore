package dblogic

import (
	"testing"
)

func TestValidateTableBlueprint(t *testing.T) {
	var err error
	tbValid := TableBlueprint{"test_table", []ColumnBlueprint{{"col_1", "TEXT"}}}
	err = validateTableBlueprint(tbValid)
	if err != nil {
		t.Fatal("Error thrown for valid TableBlueprint.")
	}

	tbNoColumnBlueprints := TableBlueprint{"test_table", nil}
	err = validateTableBlueprint(tbNoColumnBlueprints)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}

	tbNoName := TableBlueprint{"", []ColumnBlueprint{{"col 1", "TEXT"}}}
	err = validateTableBlueprint(tbNoName)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}
}

