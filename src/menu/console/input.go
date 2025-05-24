package consolemenu

import (
	"fmt"
	"formatore/src/io"
	"formatore/src/utils"
)

func (cm *ConsoleMenu) matchUserInput(s string) (string, bool) {
	specifiedFullName := func(s string) bool {
		_, ok := cm.options[s]
		return ok
	}

	if ok := specifiedFullName(s); ok {
		return s, ok
	} else {
		s, ok := cm.charsToOptionNames[s]
		return s, ok
	}
}

// Obtains user's next choice and stores it in cm.choice.
// NOTE: This function affects the choice for the next option.
func (cm *ConsoleMenu) Input() io.ResponseStatus {
	isValid := func(s string) bool {
		_, optionFullExists := cm.options[s]
		_, optionAbbrevExists := cm.charsToOptionNames[s]
		return optionFullExists || optionAbbrevExists
	}

	response, err := cm.loopUntilValidResponse(
		isValid,
	)

	utils.Assert(
		err == nil,
		fmt.Sprintf("GetStringResponse returned an unexpected error: ~%v~.", err),
	)

	if response.Okay() {
		cm.choice = response.Content
	} else {
		cm.choice = "" // Sanitize any prior choice that may be left over.
	}

	return response.Status
}
