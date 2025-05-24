package db

import (
	"database/sql"
	"fmt"
	"formatore/errors"
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

func CBsModuloAutogenCols(db *sql.DB, tableName string) ([]structs.ColumnBlueprint, error) {
	cbs, err := ColumnBlueprints(db, tableName)
	if err != nil {
		return []structs.ColumnBlueprint{}, err
	}

	if len(cbs) < 2 {
		return []structs.ColumnBlueprint{}, errors.ErrTooFewColumns
	}

	return cbs[2:], nil
}

func extractColumnNames(cbs []structs.ColumnBlueprint) []string {
	mapFn := func(cb structs.ColumnBlueprint) string {
		return cb.Name
	}
	return utils.Map(cbs, mapFn)
}

func extractColumnNamesModuloID(cbs []structs.ColumnBlueprint) ([]string, error) {
	if cbs[0].Name == "ID" {
		return []string{}, fmt.Errorf("First column name did not match 'ID'.")
	}
	return extractColumnNames(cbs)[1:], nil
}

func addTimestamp(values []string) []string {
	return append([]string{utils.UnixTimestamp()}, values...)
}

func InsertRow(db *sql.DB, tableName string, values []string) error {
	cbs, err := ColumnBlueprints(db, tableName)
	if err != nil {
		return err
	}

	valuesWithTimestamp := addTimestamp(values)
	if err := utils.ValidateAndApostrophizeValues(cbs, valuesWithTimestamp); err != nil {
		return err
	}

	columnNames, err := extractColumnNamesModuloID(cbs)
	if err != nil {
		return err
	}
	fmtColumnNames := utils.JoinWithCommasSpaces(columnNames)

	fmtValues := utils.JoinWithCommasSpaces(valuesWithTimestamp)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fmtColumnNames, fmtValues)
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert into %s: ~%v~.", tableName, err)
	}

	return nil
}
