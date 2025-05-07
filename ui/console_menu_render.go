package ui

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
			cm.display(cm.parenthesizeFirstChar(k))
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

func (cm *ConsoleMenu) RenderOnlyHeaders() {
	cm.renderEach(cm.headersPrinterFn)
}

func (cm *ConsoleMenu) Render() {
	cm.renderEach(cm.optionsPrinterFn, cm.headersPrinterFn)
}
