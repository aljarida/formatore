package menu

import (
	"formatore/utils"
	"formatore/io"
	"strings"
)

type Menu interface {
	Render() 
	Input() ResponseStatus
	Next() *Menu
}

type cmHeaders struct {
	Title string
	Guidance string
	Controls string
	Error string
}

func (cm *ConsoleMenu) Read() (string, error) {
	return cm.io.I.Read()
}

type ConsoleMenu struct {
	headers cmHeaders // UI textual displays, called on each Render().
	options map[string]func() // Functions which may be called.

	io *io.IO // IO for obtaining and sending messages.
	choice string // Stores user response after obtaining input.
	charsToOptionNames map[string]string // Example: 'q' -> "quit" => options["quit"] -> quitFn().

	next, parent *ConsoleMenu // Links to prior menu and next menu.
}

func InitConsoleMenu(cm *ConsoleMenu) *ConsoleMenu {
	cm.initCharsToOptionNames()
	utils.Assert(cm.io != nil, "IO must not be nil!")
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
