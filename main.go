package main
/*j
import (
	"database/sql"
	"formatore/ui"
	"formatore/utils"
	"fmt"
	"formatore/db"
	"strings"
)

var commonIO = &ui.IO{
	I: &ui.FmtInput{},
	O: &ui.FmtOutput{},
}

var dbase *sql.DB

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

NOTE: Preserved for reference. Must be refactored, cleaned up.
func makeTable() {
	tb, err := ui.MakeTable(commonIO)
	handleErr(err)
	err = db.CreateTable(dbase, tb)
	handleErr(err)
}

func getTableNames() string {
	names, err := db.TableNames(dbase)
	handleErr(err)
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
	handleErr(err)
	
	validator := func(s string) bool {
		return utils.Has(names, s)
	}

	commonIO.O.Display(getTableNames())

	tableName, err := commonIO.LoopUntilValidResponse(validator,
										   "Table name:",
										   "Must be an existent table.")
	handleErr(err)

	cbs, err := db.ColumnBlueprints(dbase, tableName)
	handleErr(err)
	fmt.Println(cbs)

	values, err := ui.GetValues(commonIO, cbs[2:])
	handleErr(err)
	fmt.Println(values)

	err = db.InsertRow(dbase, tableName, values)
	handleErr(err)
}


func dropAllTables() {
	err := db.DropAllTables(dbase)
	handleErr(err)
	fmt.Printf("Dropped all tables!")
}

func printTablePreview() {
	names, err := db.TableNames(dbase)
	handleErr(err)
	validator := func(s string) bool {
		return utils.Has(names, s)
	}
	tableName, err := commonIO.LoopUntilValidResponse(validator,
										   "Table name:",
										   "Must be an existent table")
	handleErr(err)
	validator = func(s string) bool {
		_, ok := utils.IsPositiveInteger(s)
		return ok
	}

	nStr, err := commonIO.LoopUntilValidResponse(validator, "Number of rows:", "Must be positive integer.")
	n, ok := utils.IsPositiveInteger(nStr)
	utils.Assert(ok == true, "Expected confirmation of a positive integer.")

	preview, err := db.PreviewLastN(dbase, tableName, n)
	handleErr(err)
	commonIO.O.Display(preview)
}

func main() {
	var err error
	dbase, err = db.ConnectToDB(enums.DBName)
	handleErr(err)
	fmt.Println("Connected to database!")

	var MainMenu *ui.ConsoleMenu

	do := true
	MainMenu = &ui.ConsoleMenu{
		IO: commonIO,
		Options: map[string]func() {
			"Create new table": makeTable,
			"Add to table": addToTable,
			"View tables": displayTableNames,
			"Drop all tables": dropAllTables,
			"Preview tables": printTablePreview,
			"Quit application" : func () { do = false },
		},
	}

	MainMenu.Initialize()

	settingsMenu := &ui.ConsoleMenu{
		IO: commonIO,
		Headers: ui.Headers{
			Title: "Header Menu",
		},
	}

	var menu *ui.ConsoleMenu
	menu = &ui.ConsoleMenu{
		PreCallback: func() { 
			commonIO.O.Display("PreCallback Working!") 
		},
		Headers: ui.Headers{
			Title: "Main v1",
			Guidance: "Quit terminal to quit!",
			Error: "ERROR: Preview build.",
		},
		Options: map[string]func() {
			"Settings": func() {
				menu.SetNext(settingsMenu)
			},
		},
	}
	menu.Initialize()

	for {
		menu.Visualize()
		menu.Input()
		next := menu.Next()
		if next != menu {
			next.Parent = menu
		}
		menu = next
	}

}
*/
