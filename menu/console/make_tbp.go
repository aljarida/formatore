package consolemenu

import (
	"formatore/io"
	"formatore/structs"
	"formatore/utils"
)

func (cm *ConsoleMenu) MakeTableBlueprint() (io.TableBlueprintResponse, error) {
	tbRes := io.TableBlueprintResponse{}
	tableNameRes, err := cm.StringResponseViaNewMenu(
		utils.IsNotReserved,
		CMHeaders{
			Title:    "=== Make Table ===",
			Guidance: "Enter a valid table name.",
			Controls: "Navigation: (q)uit -- (d)one --",
			Error:    "Invalid table name.",
		})

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
		Name:             tableNameRes.Content,
		ColumnBlueprints: questionsRes.Content,
	}
	tbRes.Status = io.InputOkay
	return tbRes, nil
}
