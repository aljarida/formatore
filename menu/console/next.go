package consolemenu

import (
	"formatore/utils"
	"formatore/enums"
)

func (cm *ConsoleMenu) Next() *ConsoleMenu {
	if utils.Has(enums.BACK_TOKENS, cm.choice) && cm.parent != nil {
		return cm.parent
	}
	
	choice, ok := cm.matchUserInput(cm.choice)
	if !ok {
		cm.headers.Error = "Choice must match available options."
		return cm
	}

	option, ok := cm.options[choice]
	if ok {
		option()
	}

	return cm.next
}
