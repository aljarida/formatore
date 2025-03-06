package main

import (
	"fmt"
	"formatore/ui"
	"time"
)

var commonIO = &ui.IO{
	I: &ui.FmtInput{},
	O: &ui.FmtOutput{},
}

func makeTable() {
	fmt.Println("Made a table!")
	time.Sleep(time.Second * 10)
}

func main () {
	var MainMenu *ui.ConsoleMenu
	var SecondaryMenu *ui.ConsoleMenu

	SecondaryMenu = &ui.ConsoleMenu{
		IO: commonIO,
		Status: "Task completed.",
		Children: map[string]*ui.ConsoleMenu{},
	}

	MainMenu = &ui.ConsoleMenu{
		IO: commonIO,
		Status: "Welcome!",
		Children: map[string]*ui.ConsoleMenu{
			"child": SecondaryMenu,
		},
		Options: map[string]func() {
			"make table": makeTable,
		},
	}

	SecondaryMenu.Children["main"] = MainMenu

	MainMenu.Initialize()
	SecondaryMenu.Initialize()

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
