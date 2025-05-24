package main

import (
	"database/sql"
	"fmt"
	"formatore/db"
	"formatore/enums"
	"formatore/io"
	"formatore/menu/console"
	"formatore/utils"
	"os"
	"strings"
)

type App struct {
	DB *sql.DB
	CM *consolemenu.ConsoleMenu
	tablesPresent bool
}

func (app *App) handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (app *App) handleQuit(res io.ResponseStatus) {
	if res == io.InputQuit {
		os.Exit(1)
	}
}

func (app *App) handleErrAndQuit(err error, res io.ResponseStatus) {
	app.handleErr(err)
	app.handleQuit(res)
}

func initializeApp() *App {
	app := &App{}

	dbase, err := db.ConnectToDB(enums.DBName)
	app.handleErr(err)

	app.DB = dbase

	cm := &consolemenu.ConsoleMenu{}
	cm.SetIO(
		&io.IO{
			I: &io.FmtInput{},
			O: &io.FmtOutput{},
		},
	)

	app.CM = cm
	consolemenu.InitConsoleMenu(cm)
	cm.SetHeaders(consolemenu.CMHeaders{
		Title:    "=== FORMATORE ===",
		Guidance: "Please choose one of the following options.",
		Controls: "Navigation: (q)uit --",
		Error:    "ERR.: Input must match available options!",
	})

	app.setMainMenuOptions()
	app.setTablesPresent()

	return app
}

func (app *App) setTablesPresent() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)
	app.tablesPresent = len(names) > 0
}

func (app *App) splashesNoTables() bool {
	if !app.tablesPresent {
		app.splash("No tables present.", "=== Failure ===")
		return true
	}
	return false
}

func (app *App) setMainMenuOptions() {
	utils.Assert(app.CM != nil, "CM must be initialized.")
	options := map[string]func(){
		"Make table":      func() { app.makeTable() },
		"Show tables":     func() { app.displayTableNames() },
		"Add entry":       func() { app.addEntryToTable() },
		"Remove table":    func() { app.removeTable() },
		"Drop all tables": func() { app.dropAllTables() },
		"Preview":         func() { app.printTablePreview() },
		"Export to CSV":   func() { app.exportToCSV() },
		"Quit Formatore":  func() { os.Exit(1) },
	}
	app.CM.SetOptions(options)
}

func (app *App) getTableName() io.StringResponse {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	tableRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Title:    "=== Tables ===",
		Guidance: "Please enter a table name.",
		Error:    "Input must be an already existing table!",
		Controls: "Navigation: (q)uit -- (b)ack --",
	}, app.tableNames())


	app.handleErrAndQuit(err, tableRes.Status)
	
	return tableRes
}

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

func (app *App) tableNames() string {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	var builder strings.Builder
	for i, n := range names {
		var fmter string
		if i == len(names) - 1 {
			fmter = "%s\n"
		} else {
			fmter = "%s"
		}
		builder.WriteString(fmt.Sprintf(fmter, n))
	}
	return builder.String()
}

func (app *App) displayTableNames() {
	if app.splashesNoTables() {
		return
	}

	app.splash(app.tableNames(), "=== Success ===")
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

	app.splash(fmt.Sprintf("Row inserted into '%s'!", tableRes.Content), "=== Success ===")
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
	
	app.splash(fmt.Sprintf("Successfully dropped table '%s'.", tableRes.Content), "=== Success ===")
	app.setTablesPresent()
}

func (app *App) dropAllTables() {
	err := db.DropAllTables(app.DB)
	app.handleErr(err)

	app.splash("All tables dropped.", "=== Success ===")

	app.tablesPresent = false
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

	app.splash(utils.JoinStrings(
		"Exported table ",
		tableRes.Content,
		" to CSV file ",
		fileName,
		".",
	), "=== Success ===")
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

func (app *App) splash(body string, title string) {
	res, err := app.CM.StringResponseViaNewMenu(
		func(s string) bool { return io.InputIsBack(s) || io.InputIsDone(s) },
		consolemenu.CMHeaders{
			Title:    title,
			Controls: "Navigation: (q)uit -- (b)ack --",
		},
		body,
	)
	app.handleErrAndQuit(err, res.Status)
}

func (app *App) loop() {
	for {
		res := app.CM.Input()
		app.handleQuit(res)
		next := app.CM.Next()
		app.CM = next
	}
}

func main() {
	app := initializeApp()
	app.loop()
}
