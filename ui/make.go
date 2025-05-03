package ui

import (
	"fmt"
	"formatore/structs"
	"formatore/utils"
)

func MakeTable(io *IO) (structs.TableBlueprint, error) {
	tableName, err := io.LoopUntilValidResponse(utils.IsNotReserved, "Table name:", "Invalid table name.")
	if err != nil {
		return structs.TableBlueprint{}, err
	}
	
	questions, err := getQuestions(io)
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
