package main

import (
	"fmt"
	"formatore/db"
	"formatore/enums"
	"formatore/ui"
	"formatore/utils"
)

func main () {
	io := ui.IO{
		I: &ui.FmtInput{},
		O: &ui.FmtOutput{},
	}
	database, err := db.ConnectToDB(enums.DBName)
	if err != nil {
		return
	}
	quit := false
	for !quit {
		io.DisplayMany(
			"(v)iew tables",
			"(c)reate a table",
			"(a)dd entry",
		)
		validator := func(s string) bool {
			return utils.Has([]string{"v", "d", "c", "a"}, s)
		}
		response, err := io.GetResponse(validator, "Response:", "Not a valid response.")
		fmt.Printf("response: %s; error: %v\n", response, err)
		if err != nil {
			return
		} else {
			switch (response) {
			case "v":
				names, err := db.TableNames(database)
				if err != nil {
					return 
				}
				for _, name := range names {
					io.O.Display(name)
				}
			case "c":
				io.O.Display("Let's create a table!")
				tb, err := ui.UiMakeTable(&io)
				if err != nil {
					return 
				}
				err = db.CreateTable(database, tb)
				if err != nil {
					fmt.Printf("Error encountered. Could not make table.\n%v", err)
					return
				}
			case "a":
				io.O.Display("Which table would you like to add to?")
			default:
				io.O.Display("Default option.")
			}
		}
	}
}
