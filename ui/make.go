package ui

import (
	"fmt"
	"formatore/structs"
	"formatore/utils"
)

func MakeTable(cm *ConsoleMenu) (structs.TableBlueprint, error) {
	tableName, err := cm.LoopUntilValidResponse(
		utils.IsNotReserved,
		cmHeaders{
			Guidance: "Table name:",
			Error: "Invalid table name."},
		)

	if err != nil {
		return structs.TableBlueprint{}, err
	}
	
	questions, err := cm.getQuestions()
	if err == ErrUserQuit { 
		return structs.TableBlueprint{}, err	
	} else if err != nil {
		return structs.TableBlueprint{}, fmt.Errorf("Error in UiMakeTable() logic: ~%v~.", err)
	}

	tb := structs.TableBlueprint{
		Name: tableName,
		ColumnBlueprints: questions,
	}
	return tb, nil
}
