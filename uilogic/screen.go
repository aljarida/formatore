package uilogic

import (
	"fmt"
)

func clearScreen() {
	fmt.Print(ansiEscapeCode)
}

func display(content string) {
	fmt.Println(content)
}

func clearThenDisplay(content string) {
	clearScreen()
	display(content)
}
