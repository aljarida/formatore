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

func (app *App) handleRes(res io.ResponseStatus) {
	if res == io.InputQuit {
		os.Exit(1)
	} else {
		utils.Assert(res == io.InputOkay, "ResponseStatus was not InputOkay.")
	}
}

func (app *App) handleErrAndRes(err error, res io.ResponseStatus) {
	app.handleErr(err)
	app.handleRes(res)
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
		"Quit Formatore": func() { os.Exit(1) },
	}
	app.CM.SetOptions(options)
}

func (app *App) makeTable() {
	tbRes, err := app.CM.MakeTableBlueprint()
	app.handleErr(err)
	// TODO: Handle non-okay responses without failing.
	utils.Assert(tbRes.Okay(), "ResponseStatus was not InputOkay.")

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
	// TODO: Pretty sure that this will need to be its own optionless menu.
	// NOTE: Good idea: Make a method in ConsoleMenu that creates a menu that
	// simply displays a splash screen and awaits "b" or "back" basically.
	app.CM.Display(app.tableNames())
}

func (app *App) addEntryToTable() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	app.displayTableNames()

	tableRes, err := app.CM.LoopUntilValidResponse(validator, consolemenu.CMHeaders{
		Guidance: "Table name:",	
		Error: "Input must be a valid table.",
		Controls: "<< (q) to quit >>",
	})
	app.handleErrAndRes(err, tableRes.Status)

	cbs, err := db.ColumnBlueprints(app.DB, tableRes.Content)
	app.handleErr(err)

	valuesRes, err := app.CM.GetValues(cbs)
	app.handleErrAndRes(err, valuesRes.Status)

	err = db.InsertRow(app.DB, tableRes.Content, valuesRes.Content)
	app.handleErr(err)
}

func (app *App) dropAllTables() {
	err := db.DropAllTables(app.DB)
	app.handleErr(err)
	// TODO: Splash screen: Dropped all tables!
	app.CM.Displayln("TEMPORARY: Dropped all tables!")
}

func (app *App) printTablePreview() {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	tableRes, err := app.CM.LoopUntilValidResponse(validator, consolemenu.CMHeaders{
		Guidance: "Table name:",
		Error: "Input must be a valid table.",
	})
	app.handleErrAndRes(err, tableRes.Status)

	validator = func(s string) bool {
		_, ok := utils.IsPositiveInteger(s)
		return ok
	}

	nStrRes, err := app.CM.LoopUntilValidResponse(validator, consolemenu.CMHeaders{
		Guidance: "Number of rows (must be positive):",
		Error: "Must be a positive integer. Try again.",
	})
	app.handleErrAndRes(err, nStrRes.Status)
	n, _ := utils.IsPositiveInteger(nStrRes.Content)

	preview, err := db.PreviewLastN(app.DB, tableRes.Content, n)
	app.handleErr(err)	

	app.CM.Display("TEMPORARY: ")
	app.CM.Displayln(preview)
}

func (app *App) loop() {
	for {
		app.CM.Render()
		app.handleRes(app.CM.Input())
		next := app.CM.Next()	
		if next != app.CM {
			app.CM.SetNext(next)
		}
	}
}

func main() {
	app := initializeApp()		
	app.loop()
}
