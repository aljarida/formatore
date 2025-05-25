package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"formatore/src/utils"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func createSiblingFolder(folderName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("Could not get program's executable path: %w.", err)
	}

	exeDir := filepath.Dir(exePath)
	siblingFolder := filepath.Join(exeDir, folderName)
	if err := os.MkdirAll(siblingFolder, 0o755); err != nil {
		return "", fmt.Errorf("Could not create the '%s' directory: %w.", folderName, err)
	}

	return siblingFolder, nil
}

func fileExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return !info.IsDir()
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func writeHeadersAndRowsToCSV(filePath string, headers []string, rows [][]string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write(headers); err != nil {
		return err
	}

	for _, row := range rows {
		if err := w.Write(row); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}

func ExportTableToCSV(db *sql.DB, tableName string, fileName string) error {
	utils.Assert(strings.HasSuffix(fileName, ".csv"), "Filename should have '.csv' extension.")

	exportFolderPath, err := createSiblingFolder("csvs")
	if err != nil {
		return err
	}

	exportFilePath := filepath.Join(exportFolderPath, fileName)
	if fileExists(exportFilePath) {
		return fmt.Errorf("File already exists at '%s'.", exportFilePath)
	}

	cbs, err := ColumnBlueprints(db, tableName)
	if err != nil {
		return err
	}

	n := len(cbs)
	utils.Assert(n > 0, "Expected more than zero columns.")

	headers := make([]string, n)
	for i, cb := range cbs {
		headers[i] = cb.Name
	}

	query := fmt.Sprintf(
		"SELECT %s FROM %s",
		utils.JoinWithCommasSpaces(headers),
		tableName,
	)

	scanFn := func(rows *sql.Rows) ([]string, error) {
		data := make([]interface{}, n)
		ptrs := make([]interface{}, n)
		for i := range data {
			ptrs[i] = &data[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		out := make([]string, n)
		for i, v := range data {
			if v == nil {
				out[i] = ""
			} else {
				out[i] = fmt.Sprintf("%v", v)
			}
		}

		return out, nil
	}

	rowData, err := queryAndScan(db, query, scanFn)
	if err != nil {
		return err
	}

	return writeHeadersAndRowsToCSV(exportFilePath, headers, rowData)
}
