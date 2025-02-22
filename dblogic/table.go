package dblogic

import (
	"database/sql"
	"fmt"
	"strings"
	"formatore/utils"
)

// Return formatted schema string for query using provided ColumnBlueprints.
func columnBlueprintsToQueryComp(cbs []utils.ColumnBlueprint) string {
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
func makeCreateQuery(tb utils.TableBlueprint) string {
	intro := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ", tb.Name)
	body := columnBlueprintsToQueryComp(tb.ColumnBlueprints)
	return utils.JoinStrings(intro,
				"(",
				utils.PKeyComp,
				utils.UnixDatetimeComp,
				body,
				");")
}

// Create a table from a provided TableBlueprint.
func CreateTable(db *sql.DB, tb utils.TableBlueprint) error {
	tb.ColumnBlueprints = formatColumnBlueprints(tb.ColumnBlueprints)
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


func scanRows[T any](
	rows *sql.Rows,
	scanFn func(*sql.Rows) (T, error)) ([]T, error) {
		var results []T
		for rows.Next() {
			item, err := scanFn(rows)
			if err != nil {
				return nil, err	
			}
			results = append(results, item)
		}
		return results, rows.Err()
}

func queryAndScan[T any](
	db *sql.DB,
	query string,
	scanFn func(*sql.Rows) (T, error)) ([]T, error) {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return scanRows(rows, scanFn)
}

// TODO: Add unit tests!
func TableNames(db *sql.DB) ([]string, error) {
	query := "SELECT name FROM sqlite_master WHERE type='table';"
	scanFn := func(rows *sql.Rows) (string, error) {
		var name string
		err := rows.Scan(&name)
		return name, err
	}
	return queryAndScan(db, query, scanFn)
}
