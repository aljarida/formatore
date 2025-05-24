package db

import (
	"database/sql"
)

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
