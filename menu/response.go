package menu

import (
	"formatore/structs"
	"formatore/utils"
	"formatore/enums"
)

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

func inputIsQuit(s string) bool {
	return utils.Has(enums.QUIT_TOKENS, s)
}

func inputIsDone(s string) bool {
	return utils.Has(enums.DONE_TOKENS, s)
}

func inputIsBack(s string) bool {
	return utils.Has(enums.BACK_TOKENS, s)
}
