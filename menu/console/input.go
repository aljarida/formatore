package consolemenu

import (
	"fmt"
	"formatore/utils"
	"formatore/io"
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
func (cm *ConsoleMenu) Input() io.ResponseStatus {
	isValid := func(s string) bool {
		_, optionFullExists := cm.options[s]
		_, optionAbbrevExists := cm.charsToOptionNames[s]
		// TODO: Strange that "InputIsBack only appears here.
		return optionFullExists || optionAbbrevExists || io.InputIsBack(s)
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
		cm.choice = response.Content
	}

	return response.Status
}
