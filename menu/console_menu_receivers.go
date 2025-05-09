package menu

import (
	"formatore/utils"
	"formatore/structs"
	"fmt"
)

func (cm *ConsoleMenu) GetValues(cbs []structs.ColumnBlueprint) (StringArrayResponse, error) {
	strArrRes := StringArrayResponse{}

	values := []string{}
	for i, cb := range cbs {
		valdator := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		headers := cmHeaders{
			Guidance: fmt.Sprintf("%d. %s (%s):", i, prettyColumnNameAsQuestion(cb.Name), cb.Type),
			Error: "Answer must match type.",
		}

		strRes, err := cm.LoopUntilValidResponse(valdator, headers)
		if err != nil {
			return strArrRes, err
		} else if strRes.Quit() {
			strArrRes.status = InputQuit
			return strArrRes, nil
		} else if strRes.Done() || strRes.UserError() { // User can not prematurely indicate "Done".
			strArrRes.status = InputUserError
			return strArrRes, nil
		} else {
			values = append(values, strRes.content)
			continue
		}
	}

	strArrRes.content = values	
	strArrRes.status = InputOkay
	return strArrRes, nil
}

// TODO: Can this function be broken up to be more DRY?
func (cm *ConsoleMenu) getQuestion() (ColumnBlueprintResponse, error) {
	cbRes := ColumnBlueprintResponse{}

	qTextResponse, err := cm.LoopUntilValidResponse(
		utils.IsNotReserved,
		cmHeaders{
			Guidance: "Question:",
			Error: "Invalid question.",
		})

	if qTextResponse.Done() || qTextResponse.Quit() {
		cbRes.status = qTextResponse.status
		return cbRes, nil
	} else if err != nil {
		return cbRes, err
	}

	qTypeResponse, err := cm.LoopUntilValidResponse(
		utils.IsValidType,
		cmHeaders{
			Guidance: "Type:",
			Error: "Invalid type.",
		})

	if qTypeResponse.Done() || qTextResponse.Quit() {
		cbRes.status = qTypeResponse.status
		return cbRes, nil
	} else if err != nil {
		return cbRes, err
	}

	cbRes.content = structs.ColumnBlueprint{
		Name: qTextResponse.content,
		Type: qTypeResponse.content,
	}
	cbRes.status = InputOkay

	return cbRes, nil
}

func (cm *ConsoleMenu) GetQuestions() (ColumnBlueprintsResponse, error) {
	cbsRes := ColumnBlueprintsResponse{}

	questions := []structs.ColumnBlueprint{}
	for {
		cbRes, err := cm.getQuestion()
		if cbRes.Done() && len(questions) > 0 {
			cbsRes.status = InputOkay
 			return cbsRes, nil
		} else if cbRes.Done() && len(questions) == 0 {
			cbsRes.status = InputUserError
			return cbsRes, nil
		} else if cbsRes.Quit() {
			cbsRes.status = InputQuit
			return cbsRes, nil
		} else if err != nil {
			return cbsRes, err
		} else {
			questions = append(questions, cbRes.content)
			continue
		}
	}
}

