package consolemenu

import (
	"formatore/structs"
	"formatore/utils"
	"formatore/io"
)

func MakeTable(cm *ConsoleMenu) (io.TableBlueprintResponse, error) {
	tbRes := io.TableBlueprintResponse{}
	tableNameRes, err := cm.LoopUntilValidResponse(
		utils.IsNotReserved,
		cmHeaders{
			Guidance: "Table name:",
			Error: "Invalid table name."},
		)

	if !tableNameRes.Okay() || err != nil {
		tbRes.Status = tableNameRes.Status
		return tbRes, err
	}
	
	questionsRes, err := cm.GetQuestions()
	if !questionsRes.Okay() || err != nil { 
		tbRes.Status = questionsRes.Status 
		return tbRes, err
	}

	tbRes.Content = structs.TableBlueprint{
		Name: tableNameRes.Content,
		ColumnBlueprints: questionsRes.Content,
	}
	tbRes.Status = io.InputOkay
	return tbRes, nil
}
