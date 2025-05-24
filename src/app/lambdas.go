package app

import (
	"fmt"
	"formatore/src/db"
	"formatore/src/menu/console"
	"formatore/src/utils"
	"strings"
)

func (app *App) makeTable() {
	tbRes, err := app.CM.MakeTableBlueprint()
	app.handleErrAndQuit(err, tbRes.Status)
	if tbRes.Back() {
		return
	}

	err = db.CreateTable(app.DB, tbRes.Content)
	app.handleErr(err)

	app.tablesPresent = true
}

func (app *App) displayTableNames() {
	if app.splashesNoTables() {
		return
	}

	app.splashSuccess(app.tableNames())
}

func (app *App) addEntryToTable() {
	if app.splashesNoTables() {
		return
	}

	tableRes := app.getTableName()
	if tableRes.Back() {
		return
	}

	cbs, err := db.CBsModuloAutogenCols(app.DB, tableRes.Content)
	app.handleErr(err)

	valuesRes, err := app.CM.GetValues(cbs)
	app.handleErrAndQuit(err, valuesRes.Status)
	if valuesRes.Back() {
		return
	}

	err = db.InsertRow(app.DB, tableRes.Content, valuesRes.Content)
	app.handleErr(err)

	app.splashSuccess(fmt.Sprintf("Row inserted into '%s'!", tableRes.Content))
}

func (app *App) removeTable() {
	if app.splashesNoTables() {
		return
	}

	tableRes := app.getTableName()
	if tableRes.Back() {
		return
	}

	err := db.DropTable(app.DB, tableRes.Content)
	app.handleErr(err)

	app.splashSuccess(fmt.Sprintf("Dropped table '%s'.", tableRes.Content))
	app.setTablesPresent()
}

func (app *App) exportToCSV() {
	if app.splashesNoTables() {
		return
	}

	tableRes := app.getTableName()
	if tableRes.Back() {
		return
	}

	validator := func(s string) bool {
		return s != "" && !strings.Contains(s, " ") && !strings.HasPrefix(s, ".")
	}

	fileNameRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Please enter a file name for the output CSV.",
		Error:    "Input must be non-empty, contain no spaces, and not be prefixed by '.'.",
	})
	app.handleErrAndQuit(err, fileNameRes.Status)
	if fileNameRes.Back() {
		return
	}

	fileName := utils.MaybeAppendDotCSV(fileNameRes.Content)

	err = db.ExportTableToCSV(app.DB, tableRes.Content, fileName)
	app.handleErr(err)

	app.splashSuccess(utils.JoinStrings(
		"Exported table ",
		tableRes.Content,
		" to CSV file ",
		fileName,
		".",
	))
}

func (app *App) printTablePreview() {
	if app.splashesNoTables() {
		return
	}

	tableRes := app.getTableName()
	if tableRes.Back() {
		return
	}

	validator := func(s string) bool {
		_, ok := utils.IsPositiveInteger(s)
		return ok
	}

	nStrRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Number of rows (must be positive):",
		Error:    "Must be a positive integer. Try again.",
		Controls: "Navigation: (q)uit -- (b)ack --",
	})
	app.handleErrAndQuit(err, nStrRes.Status)
	if nStrRes.Back() {
		return
	}

	n, _ := utils.IsPositiveInteger(nStrRes.Content)

	preview, err := db.PreviewLastN(app.DB, tableRes.Content, n)
	app.handleErr(err)

	app.splash(preview, fmt.Sprintf("=== Preview of '%s' ===", tableRes.Content))
}
