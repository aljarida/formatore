package consolemenu

import (
	"formatore/io"
	"formatore/errors"
)


// NOTE: Function only returns content with InputDone.
func (cm *ConsoleMenu) LoopUntilValidResponse(
		validator func(string) bool,
		hs cmHeaders) (io.StringResponse, error) {
	if validator == nil {
		return io.StringResponse{}, errors.ErrNeedValidator
	}

	cm.SubstituteAndRerenderOnlyHeaders(cmHeaders{Guidance: hs.Guidance})

	res := io.StringResponse{}
	for {
		input, err := cm.Read()
		if err != nil {
			return res, err
		}	

		validInput := validator(input)
		if !validInput {
			cm.SubstituteAndRerenderOnlyHeaders(cmHeaders{Error: hs.Error})
			continue
		} else if io.InputIsDone(input) { // In the event this (looping) function is called in a loop.
			res.Status = io.InputDone
			return res, nil
		} else if io.InputIsQuit(input) { // In the event user wishes to cancel behavior.
			res.Status = io.InputQuit
			return res, nil
		} else {
			res.Status = io.InputOkay
			res.Content = input
			return res, nil
		}
	}
}
