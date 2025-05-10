package consolemenu

import (
	"formatore/utils"
	"formatore/structs"
	"formatore/io"
	"fmt"
)

func (cm *ConsoleMenu) GetValues(cbs []structs.ColumnBlueprint) (io.StringArrayResponse, error) {
	strArrRes := io.StringArrayResponse{}

	values := []string{}
	for i, cb := range cbs {
		valdator := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		headers := cmHeaders{
			Guidance: fmt.Sprintf("%d. %s (%s):", i, utils.PrettyColumnNameAsQuestion(cb.Name), cb.Type),
			Error: "Answer must match type.",
		}

		strRes, err := cm.LoopUntilValidResponse(valdator, headers)
		if err != nil {
			return io.StringArrayResponse{}, err
		} else if strRes.Quit() {
			strArrRes.Status = io.InputQuit
			return strArrRes, nil
		} else if strRes.Done() || strRes.UserError() { // User can not prematurely indicate "Done".
			strArrRes.Status = io.InputUserError
			return strArrRes, nil
		} else {
			values = append(values, strRes.Content)
			continue
		}
	}

	strArrRes.Content = values	
	strArrRes.Status = io.InputOkay
	return strArrRes, nil
}

// TODO: Can this function be broken up to be more DRY?
func (cm *ConsoleMenu) getQuestion() (io.ColumnBlueprintResponse, error) {
	cbRes := io.ColumnBlueprintResponse{}

	qTextResponse, err := cm.LoopUntilValidResponse(
		utils.IsNotReserved,
		cmHeaders{
			Guidance: "Question:",
			Error: "Invalid question.",
		})

	if qTextResponse.Done() || qTextResponse.Quit() {
		cbRes.Status = qTextResponse.Status
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
		cbRes.Status = qTypeResponse.Status
		return cbRes, nil
	} else if err != nil {
		return cbRes, err
	}

	cbRes.Content = structs.ColumnBlueprint{
		Name: qTextResponse.Content,
		Type: qTypeResponse.Content,
	}
	cbRes.Status = io.InputOkay

	return cbRes, nil
}

func (cm *ConsoleMenu) GetQuestions() (io.ColumnBlueprintsResponse, error) {
	cbsRes := io.ColumnBlueprintsResponse{}

	questions := []structs.ColumnBlueprint{}
	for {
		questionRes, err := cm.getQuestion()
		if questionRes.Done() && len(questions) > 0 {
			cbsRes.Status = io.InputOkay
			cbsRes.Content = questions
 			return cbsRes, nil
		} else if questionRes.Done() && len(questions) == 0 {
			cbsRes.Status = io.InputUserError
			return cbsRes, nil
		} else if questionRes.Quit() {
			cbsRes.Status = io.InputQuit
			return cbsRes, nil
		} else if err != nil {
			return io.ColumnBlueprintsResponse{}, err
		} else {
			questions = append(questions, questionRes.Content)
			continue
		}
	}
}

