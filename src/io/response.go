package io

import (
	"formatore/src/enums"
	"formatore/src/structs"
	"formatore/src/utils"
)

type ResponseStatus int

const (
	InputOkay ResponseStatus = iota
	InputDone
	InputQuit
	InputUserError
	InputBack
)

type Response struct {
	Status ResponseStatus
}

func (r Response) Okay() bool {
	return r.Status == InputOkay
}

func (r Response) Done() bool {
	return r.Status == InputDone
}

func (r Response) Quit() bool {
	return r.Status == InputQuit
}

func (r Response) UserError() bool {
	return r.Status == InputUserError
}

func (r Response) Back() bool {
	return r.Status == InputBack
}

type StringResponse struct {
	Content string
	Response
}

type TableBlueprintResponse struct {
	Content structs.TableBlueprint
	Response
}

type ColumnBlueprintResponse struct {
	Content structs.ColumnBlueprint
	Response
}

type ColumnBlueprintsResponse struct {
	Content []structs.ColumnBlueprint
	Response
}

type StringArrayResponse struct {
	Content []string
	Response
}

func InputIsQuit(s string) bool {
	return utils.Has(enums.QUIT_TOKENS, s)
}

func InputIsDone(s string) bool {
	return utils.Has(enums.DONE_TOKENS, s)
}

func InputIsBack(s string) bool {
	return utils.Has(enums.BACK_TOKENS, s)
}
