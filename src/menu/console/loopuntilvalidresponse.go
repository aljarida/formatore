package consolemenu

import (
	"formatore/src/errors"
	"formatore/src/io"
	"formatore/src/utils"
)

const errorThreshold int = 100

// NOTE: Function only returns content with InputDone.
func (cm *ConsoleMenu) loopUntilValidResponse(validator func(string) bool) (io.StringResponse, error) {
	if validator == nil {
		return io.StringResponse{}, errors.ErrNeedValidator
	}

	cm.Render()

	counter := 0
	res := io.StringResponse{}
	for {
		input, err := cm.Read()
		if err != nil {
			return res, err
		}

		if io.InputIsDone(input) { // In the event this (looping) function is called in a loop.
			res.Status = io.InputDone
			return res, nil
		} else if io.InputIsQuit(input) { // In the event user wishes to cancel behavior.
			res.Status = io.InputQuit
			return res, nil
		} else if io.InputIsBack(input) {
			res.Status = io.InputBack
			return res, nil
		} else if validInput := validator(input); !validInput {
			counter += 1
			if counter > errorThreshold {
				return res, errors.ErrTooManyInvalidResponses
			}

			cm.SetErrorState(true)
			cm.Render()
		} else {
			cm.SetErrorState(false)
			res.Status = io.InputOkay
			res.Content = input
			return res, nil
		}
	}
}

func (cm *ConsoleMenu) StringResponseViaNewMenu(validator func(string) bool, hs CMHeaders, body ...string) (io.StringResponse, error) {
	other := InitConsoleMenu(&ConsoleMenu{
		headers: hs,
		parent:  cm,
		io:      cm.io,
		body:    utils.JoinStrArrWith(body, "\n"),
	})
	res, err := other.loopUntilValidResponse(validator)
	return res, err
}
