package db

import (
	"database/sql"
	"fmt"
	"formatore/structs"
	"formatore/utils"
)

func ColumnBlueprints(db *sql.DB, tableName string) ([]structs.ColumnBlueprint, error) {
	query := fmt.Sprintf("SELECT name, type FROM pragma_table_info('%s')", tableName)
	scanFn := func(rows *sql.Rows) (structs.ColumnBlueprint, error) {
		var colName, colType string
		err := rows.Scan(&colName, &colType)
		return structs.ColumnBlueprint{Name: colName, Type: colType}, err
	}
	return queryAndScan(db, query, scanFn)
}

func extractColumnNames(cbs []structs.ColumnBlueprint) []string {
	mapFn := func(cb structs.ColumnBlueprint) string {
		return cb.Name
	}
	return utils.Apply(cbs, mapFn)
}

func addTimestamp(values []string) []string {
	return append([]string{utils.UnixTimestamp()}, values...)
}

func InsertRow(db *sql.DB, tableName string, values []string) error {
	cbs, err := ColumnBlueprints(db, tableName)
	if err != nil {
		return err
	}

	if err := utils.ValidateAndApostrophizeValues(cbs, values); err != nil {
		return err
	}

	columnNames := extractColumnNames(cbs)
	fmtColumnNames := utils.JoinWithCommasSpaces(columnNames)

	valuesWithTimestamp := addTimestamp(values)
	fmtValues := utils.JoinWithCommasSpaces(valuesWithTimestamp)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fmtColumnNames, fmtValues)
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert into %s: ~%v~.", tableName, err)
	}

	return nil
}
