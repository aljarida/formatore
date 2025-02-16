package dblogic

import (
	"database/sql"
	"fmt"
	"strings"
)

// Return formatted schema string for query using provided ColumnBlueprints.
func columnBlueprintsToQueryComp(cbs []ColumnBlueprint) string {
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
func makeCreateQuery(tb TableBlueprint) string {
	intro := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ", tb.Name)
	body := columnBlueprintsToQueryComp(tb.ColumnBlueprints)
	return joinStrings(intro,
				"(",
				PKeyComp,
				UnixDatetimeComp,
				body,
				");")
}

// Create a table from a provided TableBlueprint.
func CreateTable(db *sql.DB, tb TableBlueprint) error {
	tb.ColumnBlueprints = formatColumnBlueprints(tb.ColumnBlueprints)
	if err := validateTableBlueprint(tb); err != nil {
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
