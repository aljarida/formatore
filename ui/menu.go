package ui

import (
	"fmt"
	"formatore/utils"
)

var BACK_TOKENS = []string{"b", "back"}

type Headers struct {
	Title string
	Guidance string
	Error string
}

type Callbacks struct {
	PreCallback func()
	PostCallback func()
}

type ConsoleMenu struct {
	IO *IO // IO for obtaining and sending messages.
	Options map[string]func() // Functions which may be called.
	Parent *ConsoleMenu // Link to prior menu.
	Headers Headers // UI textual displays.
	Callbacks Callbacks

	next *ConsoleMenu // Link to next menu.
	choice string // Stores user response.
	charToOptionName map[string]string // Example: 'q' -> "quit".
}

func (cm *ConsoleMenu) display(s string) {
	cm.IO.O.Display(s)
}

func (cm *ConsoleMenu) newline() {
	cm.displayln("")
}

func (cm *ConsoleMenu) deliverHeaders() {
	cm.displayln(cm.Headers.Title)
	cm.displayln(cm.Headers.Guidance)
	cm.displayln(cm.Headers.Error)
}

func (cm *ConsoleMenu) displayln(s string) {
	cm.IO.O.Display(fmt.Sprintf("%s\n", s))
}

// Obtains user next choice and stores it.
func (cm *ConsoleMenu) Input() {
	validator := func(s string) bool {
		_, optionFullExists := cm.Options[s]
		_, optionAbbrevExists := cm.charToOptionName[s] 	
		return optionFullExists || optionAbbrevExists || utils.Has(BACK_TOKENS, s)
	}

	response, err := cm.IO.LoopUntilValidResponse(
		validator, 
		"Choice:",
		"Choice must be a single character and match available choices.",
	)

	if err != nil {
		fmt.Printf("Error: ~%v~.", err)
	}
	
	cm.choice = response
}

func (cm *ConsoleMenu) invokePreCallback() {
	if cm.Callbacks.PreCallback != nil {
		cm.Callbacks.PreCallback()
	}
}

func (cm *ConsoleMenu) invokePostCallback() {
	if cm.Callbacks.PostCallback != nil {
		cm.Callbacks.PostCallback()
	}
}

func (cm *ConsoleMenu) deliverOptions() {
	if len(cm.Options) > 0 {
			cm.display("Options:")
			for k := range cm.Options {
				cm.display(cm.parenthesizeFirstChar(k))
			}
		}
	}

func (cm *ConsoleMenu) Visualize() {
	cm.IO.O.ClearScreen()
	cm.invokePreCallback()
	cm.deliverHeaders()
	cm.deliverOptions()
	cm.newline()
}

func (cm *ConsoleMenu) Next() *ConsoleMenu {
	if utils.Has(BACK_TOKENS, cm.choice) && cm.Parent != nil {
		return cm.Parent
	}
	
	choice, ok := cm.matchUserInput(cm.choice)
	if !ok {
		// TODO: Inform user via Error header that their choice was invalid.
		return cm
	}

	option, ok := cm.Options[choice]
	if ok {
		option()
	}

	cm.invokePostCallback()

	return cm.next
}

func (cm *ConsoleMenu) initKeyToOption() {
 	numOptions := len(cm.Options)
	keysToOptions := make(map[string]string, numOptions)

	for k := range cm.Options {
		firstChar := string(k[0])
		keysToOptions[firstChar] = k
	}

	cm.charToOptionName = keysToOptions
}

func (cm *ConsoleMenu) Initialize() {
	cm.initKeyToOption()
}

// Return a valid match if one exists.
func (cm *ConsoleMenu) matchUserInput(s string) (string, bool) {
	specifiedFullName := func(s string) bool { 
		_, ok := cm.Options[s]
		return ok
	}
	
	if ok := specifiedFullName(s); ok {
		return s, ok
	} else {
		s, ok := cm.charToOptionName[s]
		return s, ok
	}
}

type Menu interface {
	Visualize() 
	Input()
	Next() *Menu
}
