package ui

import (
	"fmt"
	"formatore/utils"
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

func (cm *ConsoleMenu) Read() (string, error) {
	return cm.io.I.Read()
}

// Obtains user's next choice and stores it in cm.choice.
func (cm *ConsoleMenu) Input() ResponseStatus {
	isValid := func(s string) bool {
		_, optionFullExists := cm.options[s]
		_, optionAbbrevExists := cm.charsToOptionNames[s]
		return optionFullExists || optionAbbrevExists || utils.Has(BACK_TOKENS, s)
	}

	response, err := cm.LoopUntilValidResponse(
		isValid, 
		cmHeaders{
			Guidance: "Action:",
			Error: "Please choose from the available options.",
		},
	)

	utils.Assert(
		err == nil, 
		fmt.Sprintf("LoopUntilValidResponse returned an unexpected error: ~%v~.", err),
	)

	if response.Okay() {
		cm.choice = response.content
	}

	return response.status
}
