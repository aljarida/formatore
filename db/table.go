package db

import (
	"database/sql"
	"fmt"
	"strings"
	"formatore/utils"
	"formatore/enums"
	"formatore/structs"
)

// Return formatted schema string for use in SQL queries.
// Uses the provided ColumnBlueprint array.
func columnBlueprintsToQueryComp(cbs []structs.ColumnBlueprint) string {
	var builder strings.Builder
	for i, cb := range cbs {
		n, t := cb.Name, cb.Type 
		fmtStr := "%s %s, "
		if i == len(cbs) - 1 {
			fmtStr = "%s %s"
		} 

		formattedComp := fmt.Sprintf(fmtStr, n, t)
		builder.WriteString(formattedComp)
	}
	return builder.String()
}

// Given a TableBlueprint, returns a formatted create statement.
// Automatically defines the primary key and the datetime fields.
// Assumes tb is a validated TableBlueprint.
func makeCreateQuery(tb structs.TableBlueprint) string {
	intro := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ", tb.Name)
	body := columnBlueprintsToQueryComp(tb.ColumnBlueprints)
	return utils.JoinStrings(intro,
				"(",
				enums.PKeyComp,
				enums.UnixDatetimeComp,
				body,
				");")
}

// Create a table from a provided TableBlueprint.
func CreateTable(db *sql.DB, tb structs.TableBlueprint) error {
	formatColumnBlueprints(tb.ColumnBlueprints)
	if err := utils.ValidateTableBlueprint(tb); err != nil {
		return err
	}

	query := makeCreateQuery(tb)
	_, err := db.Exec(query)
	return err
}

// Drop a table by its given name.
func DropTable(db *sql.DB, name string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", name)
	_, err := db.Exec(query)
	return err
}

func TableNames(db *sql.DB) ([]string, error) {
	query := "SELECT name FROM sqlite_master WHERE TYPE='table' AND name != 'sqlite_sequence';"
	scanFn := func(rows *sql.Rows) (string, error) {
		var name string
		err := rows.Scan(&name)
		return name, err
	}
	return queryAndScan(db, query, scanFn)
}

func DropAllTables(db * sql.DB) error {
	tableNames, err := TableNames(db) 
	if err != nil {
		return err
	}
	for _, name := range tableNames {
		err := DropTable(db, name)
		if err != nil {
			return err
		}
	}
	return nil
}
