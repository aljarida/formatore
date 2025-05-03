package utils

import (
	"fmt"
	"formatore/structs"
	"formatore/enums"
)

// Validates that the table name is not empty.
// Validates that the column schema is not empty.
func ValidateTableBlueprint(tb structs.TableBlueprint) error {
	if tb.Name == "" {
		return fmt.Errorf("Empty table name.")
	} else if len(tb.ColumnBlueprints) == 0 {
		return fmt.Errorf("Empty column schema.")
	}

	var err error
	err = IsValidIdentifier(tb.Name)
	if err != nil {
		return err
	}

	for _, cb := range tb.ColumnBlueprints {
		if b := IsValidType(cb.Type); !b {
			return fmt.Errorf("Invalid column type: ~%s~.", cb.Type)
		}

		if err = IsValidIdentifier(cb.Name); err != nil {
			return fmt.Errorf("At least one erroneous column name: ~%v~.", err)
		}
	}

	return nil
}

func ValidateAndApostrophizeValues(metadata []structs.ColumnBlueprint, values []string) error {
	// Metadata always has one extra column for the autoincremented key.
	metadata = metadata[1:] // Create slice that skips first entry.

	if len(metadata) != len(values) {
		return fmt.Errorf("Mismatch between number of values and number of columns")
	}

	for i := 0; i < len(metadata); i++ {
		expectedType := metadata[i].Type
		actualType := InferType(values[i])
		if expectedType != actualType {
			msg := "Type mismatch: Table column %s has type %s but value %s has type %s: ~%v~."
			return fmt.Errorf(msg, metadata[i].Name, expectedType, values[i], actualType)
		}

		// SQLite requires apostrophes around text.
		if actualType == enums.Text {
			values[i] = JoinStrings("'", values[i], "'")
		}
	}

	return nil
}
