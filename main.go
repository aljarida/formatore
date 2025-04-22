package main

import (
	"database/sql"
	"formatore/utils"
	"fmt"
	"formatore/db"
	"formatore/enums"
	"formatore/ui"
	"strings"
)

var commonIO = &ui.IO{
	I: &ui.FmtInput{},
	O: &ui.FmtOutput{},
}

var dbase *sql.DB

func mainHandleErr(err error) {
	if err != nil {
		fmt.Printf("Unexpected error: ~%v~. Terminating.", err)
	}
}

func MakeTable() {
	tb, err := ui.MakeTable(commonIO)
	mainHandleErr(err)
	db.CreateTable(dbase, tb)
}

/* NOTE: Potential callback which could be added as a property.
func ViewTables() {
	names, err := db.TableNames(dbase)
	mainHandleErr(err)
	for _, n := range names {
		commonIO.O.Display(n)
	}
	commonIO.O.Display("\n")
}
*/


func getTableNames() string {
	names, err := db.TableNames(dbase)
	mainHandleErr(err)
	var builder strings.Builder
	fmt.Printf("Found the following tables: %v", names)
	for _, n := range names { 
		builder.WriteString(fmt.Sprintf("%s\n", n))
	}
	return builder.String()
}

func addToTable() {
	names, err := db.TableNames(dbase)
	mainHandleErr(err)
	
	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	commonIO.O.Display(getTableNames())

	tableName, err := commonIO.GetResponse(validator,
										   "Table name:",
										   "Must be an existent table.")
	mainHandleErr(err)

	cbs, err := db.ColumnBlueprints(dbase, tableName)
	mainHandleErr(err)

	values, err := ui.GetValues(commonIO, cbs[2:])
	mainHandleErr(err)

	err = db.InsertRow(dbase, tableName, values)
	mainHandleErr(err)
}

func main () {
	var err error
	dbase, err = db.ConnectToDB(enums.DBName)
	mainHandleErr(err)

	var MainMenu *ui.ConsoleMenu

	MainMenu = &ui.ConsoleMenu{
		IO: commonIO,
		Status: "Welcome to Formatore.",
		Options: map[string]func() {
			"New table": MakeTable,
			"Add to table": addToTable,
		},
		// TODO: It's possibly better to create callback functions for the menus rather than creating wrapper functions. Look into this.
		Children: map[string]*ui.ConsoleMenu {
			"View tables": &ui.ConsoleMenu{IO: commonIO, Status: getTableNames()},
		},
	}

	MainMenu.Initialize()

	var menu *ui.ConsoleMenu
	menu = MainMenu

	for {
		menu.Visualize()
		menu.Input()
		next := menu.Next()
		next.Parent = menu
		menu = next
	}
}
