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
	if cm.headers.Error != "" {
		cm.Displayln(cm.headers.Error)
	}
}

func (cm *ConsoleMenu) optionsPrinterFn() {
	if len(cm.options) > 0 {
		cm.Displayln("Options:")
		for optName := range cm.options {
			cm.Displayln(utils.ParenthesizeFirstChar(optName))
		}
	}
}

func (cm *ConsoleMenu) renderEach(printerFns ...func()) {
	// TODO: Remove comment.
	// TEMPORARY COMMENTING-OUT.
	// cm.clearScreen()

	for _, printerFn := range printerFns {
		printerFn()
	}

	cm.Newline()
}

func (cm *ConsoleMenu) substituteNonEmptyHeaders(hs CMHeaders) {
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

// TODO: Does a menu which actively updates its headers need the option to not Display its options?
// TODO: Can the be reworded so it's less verbose?
func (cm *ConsoleMenu) SubstituteAndRerenderOnlyHeaders (hs CMHeaders) {
	cm.substituteNonEmptyHeaders(hs)
	cm.RenderOnlyHeaders()
}

func (cm *ConsoleMenu) RenderOnlyHeaders() {
	cm.renderEach(cm.headersPrinterFn)
}

func (cm *ConsoleMenu) Render() {
	cm.renderEach(cm.optionsPrinterFn, cm.headersPrinterFn)
}
