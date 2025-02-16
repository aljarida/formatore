package dblogic

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func getColumnMetadata(db *sql.DB, tableName string) ([]ColumnBlueprint, error) {
	query := fmt.Sprintf("SELECT name, type FROM pragma_table_info('%s')", tableName)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metadata []ColumnBlueprint
	for rows.Next() {
		var columnName, columnType string
		if err := rows.Scan(&columnName, &columnType); err != nil {
			return nil, fmt.Errorf("Error scanning table columns: ~%v~.", err)
		}
		metadata = append(metadata, ColumnBlueprint{columnName, columnType})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error with row iterator: ~%v~.", err)
	}

	return metadata, nil
}

// Infer the type of input as either INTEGER, REAL, or TEXT. (BLOB, NULL excluded.)
func InferType(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return Integer
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return Real
	}

	return Text
}

func validateAndApostrophizeValues(metadata []ColumnBlueprint, values []string) error {
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
		if actualType == Text {
			values[i] = joinStrings("'", values[i], "'")
		}
	}

	return nil
}

func getColumnNames(metadata []ColumnBlueprint) []string {
	res := make([]string, len(metadata))
	for i, metadatum := range metadata {
		res[i] = metadatum.Name
	}
	return res
}

func unixTimestamp() string {
	timestamp := time.Now().UTC().UnixNano()
	return strconv.FormatInt(timestamp, 10)
}

func addTimestamp(values []string) []string {
	return append([]string{unixTimestamp()}, values...)
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
	fmtColumnNames := joinWithCommasSpaces(columnNames)

	valuesWithTimestamp := addTimestamp(values)
	fmtValues := joinWithCommasSpaces(valuesWithTimestamp)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fmtColumnNames, fmtValues)
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert into %s: ~%v~.", tableName, err)
	}

	return nil
}
