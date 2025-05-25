package consolemenu

import (
	"fmt"
	"formatore/src/io"
	"formatore/src/structs"
	"formatore/src/utils"
	"strings"
)

func (cm *ConsoleMenu) GetValues(cbs []structs.ColumnBlueprint) (io.StringArrayResponse, error) {
	strArrRes := io.StringArrayResponse{}

	values := []string{}
	for i, cb := range cbs {
		valdator := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		headers := CMHeaders{
			Title:    "=== Values ===",
			Guidance: fmt.Sprintf("%d. %s (%s):", i, utils.PrettyColumnNameAsQuestion(cb.Name), cb.Type),
			Controls: "Navigation: (q)uit --",
			Error:    "Answer must match type.",
		}

		strRes, err := cm.StringResponseViaNewMenu(valdator, headers)
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

func (cm *ConsoleMenu) getQuestion(n int) (io.ColumnBlueprintResponse, error) {
	cbRes := io.ColumnBlueprintResponse{}
	
	textControls := "Navigation: (q)uit --"
	if n > 1 {
		textControls += " (d)one --"
	}

	qTextResponse, err := cm.StringResponseViaNewMenu(
		func(s string) bool {
			s = strings.ToUpper(s)
			return utils.IsNotReserved(s)
		},
		CMHeaders{
			Title:    fmt.Sprintf("=== Column %d Name ===", n),
			Guidance: "Please enter a column name that serves as a question.",
			Error:    "Invalid column name.",
			Controls: textControls,
		})

	if qTextResponse.Done() || qTextResponse.Quit() {
		cbRes.Status = qTextResponse.Status
		return cbRes, nil
	} else if err != nil {
		return cbRes, err
	}

	qTypeResponse, err := cm.StringResponseViaNewMenu(
		utils.IsValidType,
		CMHeaders{
			Title:    fmt.Sprintf("=== Column %d Type ===", n),
			Guidance: "Please enter a type for the question (TEXT, INTEGER, REAL)",
			Error:    "Invalid type.",
			Controls: "Navigation: (q)uit --",
		})

	if qTypeResponse.Done() || qTypeResponse.Quit() {
		cbRes.Status = qTypeResponse.Status
		return cbRes, nil
	} else if err != nil {
		return cbRes, err
	}

	qTypeResponse.Content = strings.ToUpper(qTypeResponse.Content)

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
	n := 1
	for {
		questionRes, err := cm.getQuestion(n)
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
			n += 1
			continue
		}
	}
}
