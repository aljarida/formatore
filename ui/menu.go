package ui

import (
	"fmt"
	"formatore/utils"
)

var BACK_TOKENS = []string{"b", "back"}

type ConsoleMenu struct {
	IO *IO // Displays status and options.
	Status string // For displaying main text.
	Options map[string]func() // Functions which may be called.
	Children map[string]*ConsoleMenu // Menus which may be navigated to.
	Parent *ConsoleMenu // Link to prior menu.

	choice string // Stores user response.
	keyToOption map[string]string // Example: 'q' -> "quit".
	keyToChild map[string]string // Example: 'n' -> "next".
}

func (cm *ConsoleMenu) display(s string) {
	cm.IO.O.Display(s)
}

// Serves as a wrapper for the output's Display method.
func (cm *ConsoleMenu) displayln(s string) {
	cm.IO.O.Display(fmt.Sprintf("%s\n", s))
}

// Takes a string, returns parentheses around the first character.
func (cm *ConsoleMenu) pretty(s string) string {
	if len(s) < 1 {
		return s
	}

	extra := ""
	if len(s) > 1 {
		extra = s[1:]
	}

	return utils.JoinStrings("(", string(s[0]), ")", extra)
}

// Obtains user next choice and stores it.
func (cm *ConsoleMenu) Input() {
	validator := func(s string) bool {
		_, childFullExists := cm.Children[s]
		_, optionFullExists := cm.Options[s]
		_, childAbbrevExists := cm.keyToChild[s]
		_, optionAbbrevExists := cm.keyToOption[s] 	

		existsFull := childFullExists || optionFullExists
		existsAbbrev := childAbbrevExists || optionAbbrevExists

		return existsAbbrev || existsFull || utils.Has(BACK_TOKENS, s)
	}

	response, err := cm.IO.GetResponse(
		validator, 
		"Choice:",
		"Choice must be a single character and match available choices.",
	)

	if err != nil {
		fmt.Printf("Error: ~%v~.", err)
	}
	
	cm.choice = response
}

func (cm *ConsoleMenu) Visualize() {
	cm.IO.O.ClearScreen()

	cm.displayln(cm.Status)

	if len(cm.Options) > 0 {
		cm.display("Options:")
		for k := range cm.Options {
			cm.display(cm.pretty(k))
		}
	}
	cm.display("\n")

	if len(cm.Children) > 0 {
		cm.display("Menus:")
		for k := range cm.Children {
			cm.display(cm.pretty(k))
		}
	}
	cm.display("\n")
}

// Matches user input to correct command.
// If the user chose an option, returns self.
// If the user chose a child, returns appropriate child.
// If the user chose back, returns the parent.
func (cm *ConsoleMenu) Next() *ConsoleMenu {
	if utils.Has(BACK_TOKENS, cm.choice) && cm.Parent != nil {
		return cm.Parent
	}
	
	choice := cm.match(cm.choice)
	child, ok := cm.Children[choice]
	if ok {
		return child
	}

	option, ok := cm.Options[choice]
	if ok {
		option()
	}

	return cm
}

func (cm *ConsoleMenu) initKeyToChild() {
	numChildren := len(cm.Children)
	keysToChildren := make(map[string]string, numChildren)
	
	for k := range cm.Children {
		firstChar := string(k[0])
		keysToChildren[firstChar] = k
	}

	cm.keyToChild = keysToChildren
}

func (cm *ConsoleMenu) initKeyToOption() {
 	numOptions := len(cm.Options)
	keysToOptions := make(map[string]string, numOptions)

	for k := range cm.Options {
		firstChar := string(k[0])
		keysToOptions[firstChar] = k
	}

	cm.keyToOption = keysToOptions
}

// Logic to intialize necessary dictionaries for matching.
func (cm *ConsoleMenu) Initialize() {
	cm.initKeyToChild()
	cm.initKeyToOption()
}

// Return a valid match if one exists.
func (cm *ConsoleMenu) match(s string) string {
	match, ok := cm.keyToChild[s]
	if ok {
		return match
	}

	match, ok = cm.keyToOption[s]
	if ok {
		return match
	}

	return s
}

type Menu interface {
	Visualize() 
	Input()
	Next() *Menu
}
