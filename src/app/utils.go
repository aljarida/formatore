package app

import (
	"fmt"
	"formatore/src/db"
	"formatore/src/io"
	"formatore/src/menu/console"
	"formatore/src/utils"
	"os"
	"strings"
)

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

func (app *App) tableNames() string {
	names, err := db.TableNames(app.DB)
	app.handleErr(err)

	var builder strings.Builder
	for i, n := range names {
		var fmter string
		if i == len(names)-1 {
			fmter = "%s\n"
		} else {
			fmter = "%s"
		}
		builder.WriteString(fmt.Sprintf(fmter, n))
	}
	return builder.String()
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

func (app *App) splashesNoTables() bool {
	if !app.tablesPresent {
		app.splash("No tables present.", "=== Failure ===")
		return true
	}
	return false
}

func (app *App) splashSuccess(s string) {
	app.splash(s, "=== Success ===")
}
