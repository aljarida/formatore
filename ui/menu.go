package ui

import (
	"strings"
)

// TODO: Determine where these BACK_TOKENS should go. Utils?
var BACK_TOKENS = []string{"b", "back"}

type Menu interface {
	Render() 
	Input()
	Next() *Menu
}

type cmHeaders struct {
	Title string
	Guidance string
	Controls string
	Error string
}

type ConsoleMenu struct {
	headers cmHeaders // UI textual displays, called on each Render().
	options map[string]func() // Functions which may be called.

	io *IO // IO for obtaining and sending messages.
	choice string // Stores user response after obtaining input.
	charsToOptionNames map[string]string // Example: 'q' -> "quit" => options["quit"] -> quitFn().

	next *ConsoleMenu // Link to any next menu or self.
	parent *ConsoleMenu // Link to prior menu.
}

func InitConsoleMenu(cm *ConsoleMenu) *ConsoleMenu {
	cm.initCharsToOptionNames()
	return cm
}

func (cm *ConsoleMenu) initCharsToOptionNames() {
 	numOptions := len(cm.options)
	charsToOptionNames := make(map[string]string, numOptions)

	for k := range cm.options {
		firstChar := string(k[0])

		firstCharUpper := strings.ToUpper(firstChar)
		charsToOptionNames[firstCharUpper] = k

		firstCharLower := strings.ToLower(firstChar)
		charsToOptionNames[firstCharLower] = k
	}

	cm.charsToOptionNames = charsToOptionNames
}
