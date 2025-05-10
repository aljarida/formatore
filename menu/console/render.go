package consolemenu

import (
	"formatore/utils"
)

func (cm *ConsoleMenu) display(s string) {
	cm.io.O.Display(s)
}

func (cm *ConsoleMenu) displayln(s string) {
	cm.display(s + "\n")
}

func (cm *ConsoleMenu) newline() {
	cm.displayln("")
}

func (cm *ConsoleMenu) clearScreen() {
	cm.io.O.ClearScreen()
}


func (cm *ConsoleMenu) headersPrinterFn() {
	if cm.headers.Title != "" {
		cm.displayln(cm.headers.Title)
	}
	if cm.headers.Guidance != "" {
		cm.displayln(cm.headers.Guidance)
	}
	if cm.headers.Controls != "" {
		cm.displayln(cm.headers.Controls)
	}
	if cm.headers.Error != "" {
		cm.displayln(cm.headers.Error)
	}
}

func (cm *ConsoleMenu) optionsPrinterFn() {
	if len(cm.options) > 0 {
		cm.display("Options:")
		for k := range cm.options {
			cm.display(utils.ParenthesizeFirstChar(k))
		}
	}
}

func (cm *ConsoleMenu) renderEach(printerFns ...func()) {
	cm.clearScreen()

	for _, printerFn := range printerFns {
		printerFn()
	}

	cm.newline()
}

func (cm *ConsoleMenu) substituteNonEmptyHeaders(hs cmHeaders) {
	const emptyString string = ""
	if hs.Title != emptyString {
		cm.SetHeaderTitle(hs.Title)
	}
	if hs.Guidance != emptyString {
		cm.SetHeaderGuidance(hs.Guidance)
	}
	if hs.Error  != emptyString {
		cm.SetHeaderError(hs.Error)
	}
	if hs.Controls != emptyString {
		cm.SetHeaderControls(hs.Controls)
	}
}

// TODO: Does a menu which actively updates its headers need the option to not display its options?
// TODO: Can the be reworded so it's less verbose?
func (cm *ConsoleMenu) SubstituteAndRerenderOnlyHeaders (hs cmHeaders) {
	cm.substituteNonEmptyHeaders(hs)
	cm.RenderOnlyHeaders()
}

func (cm *ConsoleMenu) RenderOnlyHeaders() {
	cm.renderEach(cm.headersPrinterFn)
}

func (cm *ConsoleMenu) Render() {
	cm.renderEach(cm.optionsPrinterFn, cm.headersPrinterFn)
}
