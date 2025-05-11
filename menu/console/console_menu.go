package consolemenu

import (
	"formatore/utils"
	"formatore/io"
	"strings"
)

func (cm *ConsoleMenu) Read() (string, error) {
	return cm.io.I.Read()
}

type CMHeaders struct {
	Title string
	Guidance string
	Controls string
	Error string
}

type ConsoleMenu struct {
	headers CMHeaders // UI textual displays, called on each Render().
	body string // Primary text.
	options map[string]func() // Functions which may be called.
	charsToOptionNames map[string]string // Example: 'q' -> "quit" => options["quit"] -> quitFn().

	io *io.IO // IO for obtaining and sending messages.
	choice string // Stores user response after obtaining input.
	errorState bool

	next, parent *ConsoleMenu // Links to prior menu and next menu.
}

func InitConsoleMenu(cm *ConsoleMenu) *ConsoleMenu {
	utils.Assert(cm.options == nil, "Can not initialize an already initialized CM!")
	cm.options = make(map[string]func(), 8)
	utils.Assert(cm.io != nil, "IO must not be nil!")
	return cm
}

func (cm *ConsoleMenu) initCharsToOptionNames() {
 	numOptions := len(cm.options)
	charsToOptionNames := make(map[string]string, numOptions)

	for optName := range cm.options {
		firstChar := string(optName[0])
		_, ok := charsToOptionNames[firstChar]
		utils.Assert(ok == false, "There can not be two options of the same first character!")

		firstCharUpper := strings.ToUpper(firstChar)
		charsToOptionNames[firstCharUpper] = optName

		firstCharLower := strings.ToLower(firstChar)
		charsToOptionNames[firstCharLower] = optName
	}

	cm.charsToOptionNames = charsToOptionNames
}
