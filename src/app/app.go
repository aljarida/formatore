package app

import (
	"database/sql"
	"formatore/src/db"
	"formatore/src/enums"
	"formatore/src/io"
	"formatore/src/utils"
	"formatore/src/menu/console"
	"os"
)

type App struct {
	DB *sql.DB
	CM *consolemenu.ConsoleMenu
	tablesPresent bool
}

func (app *App) setMainMenuOptions() {
	utils.Assert(app.CM != nil, "CM must be initialized.")
	options := map[string]func(){
		"Make table":      func() { app.makeTable() },
		"Show tables":     func() { app.displayTableNames() },
		"Add entry to table":       func() { app.addEntryToTable() },
		"Remove table":    func() { app.removeTable() },
		"Preview table":         func() { app.printTablePreview() },
		"Export table to CSV":   func() { app.exportToCSV() },
		"Quit":  func() { os.Exit(1) },
	}
	app.CM.SetOptions(options)
}

func (app *App) setTablesPresent() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)
	app.tablesPresent = len(names) > 0
}

func NewApp() *App {
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
		Error:    "Input must match available options!",
	})

	app.setMainMenuOptions()
	app.setTablesPresent()

	return app
}

func (app *App) Loop() {
	for {
		res := app.CM.Input()
		app.handleQuit(res)
		next := app.CM.Next()
		app.CM = next
	}
}
