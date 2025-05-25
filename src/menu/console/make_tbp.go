package consolemenu

import (
	"formatore/src/io"
	"formatore/src/structs"
	"formatore/src/utils"
)

func (cm *ConsoleMenu) MakeTableBlueprint(invalidNames ...string) (io.TableBlueprintResponse, error) {
	tbRes := io.TableBlueprintResponse{}

	validator := func(s string) bool { return utils.IsNotReserved(s) && !utils.Has(invalidNames, s) }
	tableNameRes, err := cm.StringResponseViaNewMenu(
		validator,
		CMHeaders{
			Title:    "=== Make Table ===",
			Guidance: "Enter a valid table name.",
			Controls: "Navigation: (q)uit -- (b)ack --",
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
