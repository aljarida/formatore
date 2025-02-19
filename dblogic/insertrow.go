package dblogic

import (
	"database/sql"
	"fmt"
	"formatore/utils"
)

func getColumnMetadata(db *sql.DB, tableName string) ([]utils.ColumnBlueprint, error) {
	query := fmt.Sprintf("SELECT name, type FROM pragma_table_info('%s')", tableName)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metadata []utils.ColumnBlueprint
	for rows.Next() {
		var columnName, columnType string
		if err := rows.Scan(&columnName, &columnType); err != nil {
			return nil, fmt.Errorf("Error scanning table columns: ~%v~.", err)
		}
		metadata = append(metadata, utils.ColumnBlueprint{columnName, columnType})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error with row iterator: ~%v~.", err)
	}

	return metadata, nil
}

func validateAndApostrophizeValues(metadata []utils.ColumnBlueprint, values []string) error {
	// Metadata always has one extra column for the autoincremented key.
	metadata = metadata[1:] // Create slice that skips first entry.

	if len(metadata) != len(values) {
		return fmt.Errorf("Mismatch between number of values and number of columns")
	}

	for i := 0; i < len(metadata); i++ {
		expectedType := metadata[i].Type
		actualType := utils.InferType(values[i])
		if expectedType != actualType {
			msg := "Type mismatch: Table column %s has type %s but value %s has type %s: ~%v~."
			return fmt.Errorf(msg, metadata[i].Name, expectedType, values[i], actualType)
		}

		// SQLite requires apostrophes around text.
		if actualType == utils.Text {
			values[i] = utils.JoinStrings("'", values[i], "'")
		}
	}

	return nil
}

func getColumnNames(metadata []utils.ColumnBlueprint) []string {
	res := make([]string, len(metadata))
	for i, metadatum := range metadata {
		res[i] = metadatum.Name
	}
	return res
}


func addTimestamp(values []string) []string {
	return append([]string{utils.UnixTimestamp()}, values...)
}

func InsertRow(db *sql.DB, tableName string, values []string) error {
	metadata, err := getColumnMetadata(db, tableName)
	if err != nil {
		return err
	}

	if err := validateAndApostrophizeValues(metadata, values); err != nil {
		return err
	}

	columnNames := getColumnNames(metadata)
	fmtColumnNames := utils.JoinWithCommasSpaces(columnNames)

	valuesWithTimestamp := addTimestamp(values)
	fmtValues := utils.JoinWithCommasSpaces(valuesWithTimestamp)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fmtColumnNames, fmtValues)
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert into %s: ~%v~.", tableName, err)
	}

	return nil
}
