package ui

import (
	"formatore/utils"
	"formatore/structs"
	"fmt"
)

type IO struct {
	I InputReader
	O OutputDisplay
}

type ResponseStatus int

const (
	InputOkay ResponseStatus = iota
	InputDone
	InputQuit
	InputUserError
)

type Response struct {
	status ResponseStatus
}

func (r Response) Okay() bool {
	return r.status == InputOkay
}

func (r Response) Done() bool {
	return r.status == InputDone
}

func (r Response) Quit() bool {
	return r.status == InputQuit
}

func (r Response) UserError() bool {
	return r.status == InputUserError
}

type StringResponse struct {
	content string
	Response
}

type ColumnBlueprintResponse struct {
	content structs.ColumnBlueprint
	Response
}

type ColumnBlueprintsResponse struct {
	content []structs.ColumnBlueprint
	Response
}

type StringArrayResponse struct {
	content []string
	Response
}

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

func (cm *ConsoleMenu) LoopUntilValidResponse(
		validator func(string) bool,
		hs cmHeaders) (StringResponse, error) {
	res := StringResponse{}
	if validator == nil {
		return res, ErrNeedValidator
	}

	cm.SubstituteAndRerenderOnlyHeaders(cmHeaders{Guidance: hs.Guidance})
	for {
		input, err := cm.Read()
		if err != nil {
			return res, err
		}	

		validInput := validator(input)
		if !validInput {
			cm.SubstituteAndRerenderOnlyHeaders(cmHeaders{Error: hs.Error})
			continue
		} else if isDone(input) { // In the event this (looping) function is called in a loop.
			res.status = InputDone
			return res, nil
		} else if isQuit(input) { // In the event user wishes to cancel behavior.
			res.status = InputQuit
			return res, nil
		} else {
			res.status = InputOkay
			res.content = input
			return res, nil
		}
	}
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

func (cm *ConsoleMenu) getQuestions() (ColumnBlueprintsResponse, error) {
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

