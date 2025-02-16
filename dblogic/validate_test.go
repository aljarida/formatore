package dblogic

import (
	"testing"
	"formatore/structs"
)

func TestValidateTableBlueprint(t *testing.T) {
	var err error
	tbValid := structs.TableBlueprint{"test_table", []structs.ColumnBlueprint{{"col_1", "TEXT"}}}
	err = validateTableBlueprint(tbValid)
	if err != nil {
		t.Fatal("Error thrown for valid TableBlueprint.")
	}

	tbNoColumnBlueprints := structs.TableBlueprint{"test_table", nil}
	err = validateTableBlueprint(tbNoColumnBlueprints)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}

	tbNoName := structs.TableBlueprint{"", []structs.ColumnBlueprint{{"col 1", "TEXT"}}}
	err = validateTableBlueprint(tbNoName)
	if err == nil {
		t.Fatal("No error for invalid TableBlueprint.")
	}
}

