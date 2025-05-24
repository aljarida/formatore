package main

import (
	"database/sql"
	"formatore/utils"
	"formatore/db"
	"formatore/io"
	"formatore/enums"
	"formatore/menu/console"
	"strings"
	"fmt"
	"os"
)

type App struct {
	DB *sql.DB
	CM *consolemenu.ConsoleMenu
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
		Title: "=== Formatore ===",
		Guidance: "Please choose one of the following options.",
		Controls: "(q) to quit.",
		Error: "ERR.: Input must match available options!",
	})

	app.setMainMenuOptions()

	return app
}

func (app *App) setMainMenuOptions() {
	utils.Assert(app.CM != nil, "CM must be initialized.")
	options := map[string]func() {
		"Make table": func() { app.makeTable() },
		"Show tables": func() { app.displayTableNames() },
		"Add entry": func() { app.addEntryToTable() },
		"Drop all tables": func() { app.dropAllTables() },
		"Preview": func() { app.printTablePreview() },
		"Export to CSV": func() { app.exportToCSV() },
		"Quit Formatore": func() { os.Exit(1) },
		"Testing!": func() { app.splash("Testing well?!") },
	}
	app.CM.SetOptions(options)
}

func (app *App) makeTable() {
	tbRes, err := app.CM.MakeTableBlueprint()
	app.handleErrAndQuit(err, tbRes.Status)
	if tbRes.Back() {
		return
	}

	err = db.CreateTable(app.DB, tbRes.Content)
	app.handleErr(err)
}

func (app *App) tableNames() string {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)
	
	var builder strings.Builder
	for i, n := range names { 
		builder.WriteString(fmt.Sprintf("%d: %s\n", i + 1, n))
	}
	return builder.String()
}

func (app *App) displayTableNames() {
	app.splash(app.tableNames())
}

func (app *App) addEntryToTable() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)
	if len(names) == 0 {
		app.splash("No tables to add to.")
		return
	}

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	tableRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Title: "=== Table title ===",
		Guidance: "Please enter a table name",	
		Error: "Input must be an already existing table!",
	}, app.tableNames())

	app.handleErrAndQuit(err, tableRes.Status)
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

	app.splash("Row successfully inserted!")
}

func (app *App) dropAllTables() {
	err := db.DropAllTables(app.DB)
	app.handleErr(err)

	app.splash("All tables successfully dropped!")
}

func (app *App) exportToCSV() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	tableRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Please enter a table name.",
		Error: "Input must be a valid table.",
	}, app.tableNames())

	app.handleErrAndQuit(err, tableRes.Status)
	if tableRes.Back() {
		return
	}

	validator = func(s string) bool {
		return s != "" && !strings.Contains(s, " ") && !strings.HasPrefix(s, ".")
	}

	fileNameRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Please enter a file name for the output CSV.",
		Error: "Input must be non-empty, contain no spaces, and not be prefixed by '.'.",
	})
	app.handleErrAndQuit(err, fileNameRes.Status)
	if fileNameRes.Back() {
		return
	}

	fileName := utils.MaybeAppendDotCSV(fileNameRes.Content)
	
	err = db.ExportTableToCSV(app.DB, tableRes.Content, fileName)
	app.handleErr(err)

	app.splash(utils.JoinStrings(
		"Successfully exported table ",
		tableRes.Content,
		" to CSV file ",
		fileName,
		".",
	))
}

func (app *App) printTablePreview() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	tableRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Please enter a table name.",
		Error: "Input must be a valid table.",
	}, app.tableNames())

	app.handleErrAndQuit(err, tableRes.Status)
	if tableRes.Back() {
		return
	}

	validator = func(s string) bool {
		_, ok := utils.IsPositiveInteger(s)
		return ok
	}

	nStrRes, err := app.CM.StringResponseViaNewMenu(validator, consolemenu.CMHeaders{
		Guidance: "Number of rows (must be positive):",
		Error: "Must be a positive integer. Try again.",
	})
	app.handleErrAndQuit(err, nStrRes.Status)
	if nStrRes.Back() {
		return
	}

	n, _ := utils.IsPositiveInteger(nStrRes.Content)

	preview, err := db.PreviewLastN(app.DB, tableRes.Content, n)
	app.handleErr(err)	

	app.splash(preview)
}

// TODO: Modify this function to accept a title string.
func (app *App) splash(body string) {
	res, err := app.CM.StringResponseViaNewMenu(
		func(s string) bool { return io.InputIsBack(s) || io.InputIsDone(s) },
		consolemenu.CMHeaders{
			Title: "=== Splash Screen ===",
			Controls: "(d) or (b) to go back",
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
