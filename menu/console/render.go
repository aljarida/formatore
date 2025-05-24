package consolemenu

import (
	"formatore/utils"
)

func (cm *ConsoleMenu) Display(s string) {
	cm.io.O.Display(s)
}

func (cm *ConsoleMenu) Displayln(s string) {
	cm.Display(s + "\n")
}

func (cm *ConsoleMenu) Newline() {
	cm.Displayln("")
}

func (cm *ConsoleMenu) clearScreen() {
	cm.io.O.ClearScreen()
}

func (cm *ConsoleMenu) headersPrinterFn() {
	if cm.headers.Title != "" {
		cm.Displayln(cm.headers.Title)
	}
	if cm.headers.Guidance != "" {
		cm.Displayln(cm.headers.Guidance)
	}
	if cm.headers.Controls != "" {
		cm.Displayln(cm.headers.Controls)
	}
	if cm.headers.Error != "" && cm.errorState {
		cm.Displayln(cm.headers.Error)
	}
}

func (cm *ConsoleMenu) optionsPrinterFn() {
	if len(cm.options) == 0 {
		return
	}

	sortedKeys := utils.GetSortedKeys(cm.options)

	cm.Displayln("Options:")
	for _, optName := range sortedKeys {
		cm.Displayln(utils.ParenthesizeFirstChar(optName))
	}
}

func (cm *ConsoleMenu) renderEach(printerFns ...func()) {
	cm.clearScreen()

	for _, printerFn := range printerFns {
		printerFn()
	}
}

func (cm *ConsoleMenu) bodyPrinterFn() {
	if cm.body != "" {
		cm.Displayln(cm.body)
	}
}

func (cm *ConsoleMenu) Render() {
	cm.renderEach(cm.headersPrinterFn, cm.optionsPrinterFn, cm.bodyPrinterFn)
}
