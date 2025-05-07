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

func (cm *ConsoleMenu) GetValues(cbs []structs.ColumnBlueprint) ([]string, error) {
	values := []string{}
	for i, cb := range cbs {
		valdator := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		headers := cmHeaders{
			Guidance: fmt.Sprintf("%d. %s (%s):", i, prettyColumnNameAsQuestion(cb.Name), cb.Type),
			Error: "Answer must match type.",
		}

		res, err := cm.LoopUntilValidResponse(valdator, headers)
		if err != nil {
			// TODO: Handle error.		
			// Error is ErrUserQuit?
			// Error is ErrUserDone?
			// Error is something else?
		}

		values = append(values, res)
	}

	return values, nil
}

func (cm *ConsoleMenu) LoopUntilValidResponse(
		validator func(string) bool,
		hs cmHeaders) (string, error) {
	if validator == nil {
		return "", ErrNeedValidator
	}

	cm.SetHeaderGuidance(hs.Guidance)
	cm.RenderOnlyHeaders()
	for {
		response, err := cm.Read()
		if err != nil {
			return "", err
		}	

		valid := validator(response)
		if !valid {
			cm.SetHeaderError(hs.Error)
			cm.RenderOnlyHeaders()
			continue
		} else if isDone(response) { // In the event this function is called in a loop.
			return "", nil
		} else if isQuit(response) {
			return "", ErrUserQuit
		} else {
			return response, nil
		}
	}
}

func (cm *ConsoleMenu) getQuestion() (structs.ColumnBlueprint, error) {
	qText, err := cm.LoopUntilValidResponse(
		utils.IsNotReserved,
		Headers{
			Guidance: "Question:",
			Error: "Invalid question.",
		})

	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	qType, err := cm.LoopUntilValidResponse(
		utils.IsValidType,
		Headers{
			Guidance: "Type:",
			Error: "Invalid type.",
		})

	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	return structs.ColumnBlueprint{Name: qText, Type: qType}, nil
}

func (cm *ConsoleMenu) getQuestions() ([]structs.ColumnBlueprint, error) {
	questions := []structs.ColumnBlueprint{}
	for {
		question, err := cm.getQuestion()
		if err == ErrUserDone && len(questions) > 0 {
 			return questions, nil
		} else if err == ErrUserDone { // && len(questions == 0)
			return []structs.ColumnBlueprint{}, ErrUserQuit
		} else if err != nil {
			// TODO: Handle error appropriately!
			// TODO: Handle error.		
			// Error is ErrUserQuit?
			// Error is ErrUserDone?
			// Error is something else?
			return []structs.ColumnBlueprint{}, err
		} else {
			questions = append(questions, question)
		}
	}
}

