package dblogic

import (
	"database/sql"
	"fmt"
	"formatore/utils"
)

func getColumnMetadata(db *sql.DB, tableName string) ([]utils.ColumnBlueprint, error) {
	query := fmt.Sprintf("SELECT name, type FROM pragma_table_info('%s')", tableName)
	scanFn := func(rows *sql.Rows) (utils.ColumnBlueprint, error) {
		var colName, colType string
		err := rows.Scan(&colName, &colType)
		return utils.ColumnBlueprint{Name: colName, Type: colType}, err
	}
	return queryAndScan(db, query, scanFn)
}

// TODO: Move this into the utils.
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
