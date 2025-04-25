package main

import (
	"database/sql"
	"formatore/utils"
	"strconv"
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
		panic(err)
	}
}

func MakeTable() {
	tb, err := ui.MakeTable(commonIO)
	mainHandleErr(err)
	err = db.CreateTable(dbase, tb)
	mainHandleErr(err)
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
	for _, n := range names { 
		builder.WriteString(fmt.Sprintf("%s\n", n))
	}
	return builder.String()
}

func displayTableNames() {
	commonIO.O.Display(getTableNames())
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
	fmt.Println(cbs)

	values, err := ui.GetValues(commonIO, cbs[2:])
	mainHandleErr(err)
	fmt.Println(values)

	err = db.InsertRow(dbase, tableName, values)
	mainHandleErr(err)
}


func dropAllTables() {
	err := db.DropAllTables(dbase)
	mainHandleErr(err)
	fmt.Printf("Dropped all tables!")
}

// TODO: Move to utils. Check that this wasn't implemented already.
func isPositiveInteger(s string) (bool, int) {
    n, err := strconv.Atoi(s)
    if err != nil || n <= 0 {
        return false, 0
    }
    return true, n
}

func printTablePreview() {
	names, err := db.TableNames(dbase)
	mainHandleErr(err)
	validator := func(s string) bool {
		return utils.Has(names, s)
	}
	tableName, err := commonIO.GetResponse(validator,
										   "Table name:",
										   "Must be an existent table")
	mainHandleErr(err)
	validator = func(s string) bool {
		ok, _ := isPositiveInteger(s)
		return ok
	}
	nStr, err := commonIO.GetResponse(validator, "Number of rows:", "Must be positive integer.")
	_, n := isPositiveInteger(nStr)
	preview, err := db.PreviewLastN(dbase, tableName, n)
	mainHandleErr(err)
	commonIO.O.Display(preview)
}

func main () {
	var err error
	dbase, err = db.ConnectToDB(enums.DBName)
	mainHandleErr(err)
	fmt.Println("Connected to database!")

	var MainMenu *ui.ConsoleMenu

	MainMenu = &ui.ConsoleMenu{
		IO: commonIO,
		InitialMessage: "Welcome to Formatore.",
		Options: map[string]func() {
			"New table": MakeTable,
			"Add to table": addToTable,
			"View tables": displayTableNames,
			"Drop all tables!": dropAllTables,
			"Preview tables!": printTablePreview,
		},
	}

	MainMenu.Initialize()
	fmt.Println("Initialized main menu!")

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
