package dblogic

import (
	"fmt"
	"formatore/structs"
)

// Validates that the table name is not empty.
// Validates that the column schema is not empty.
func validateTableBlueprint(tb structs.TableBlueprint) error {
	if tb.Name == "" {
		return fmt.Errorf("Empty table name.")
	} else if len(tb.ColumnBlueprints) == 0 {
		return fmt.Errorf("Empty column schema.")
	}

	var err error
	err = IsValidIdentifer(tb.Name)
	if err != nil {
		return err
	}

	for _, cb := range tb.ColumnBlueprints {
		if b := IsValidType(cb.Type); !b {
			return fmt.Errorf("Invalid column type: ~%s~.", cb.Type)
		}

		if err = IsValidIdentifer(cb.Name); err != nil {
			return fmt.Errorf("At least one erroneous column name: ~%v~.", err)
		}
	}

	return nil
}
