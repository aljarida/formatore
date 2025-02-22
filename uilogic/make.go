package uilogic

import (
	"fmt"
	"formatore/utils"
)

func UiMakeTable(io *IO) (utils.TableBlueprint, error) {
	tableName, err := io.getResponse(utils.IsNotReserved, "Table name:", "Invalid table name.")
	if err != nil {
		return utils.TableBlueprint{}, err
	}
	
	questions, err := getQuestions(io)
	if err == ErrUserQuit { 
		return utils.TableBlueprint{}, err	
	} else if err != nil {
		return utils.TableBlueprint{}, fmt.Errorf("Error in UiMakeTable() logic: ~%v~.", err)
	}

	tb := utils.TableBlueprint{
		Name: tableName,
		ColumnBlueprints: questions,
	}
	return tb, nil
}
