package res

import "errors"

type Errs interface {
	GetCode() int
	GetShowMsg() string
	error
}

type err struct {
	code    int
	showMsg string
	error
}

func New(code int, msg, showMsg string) Errs {
	e := &err{
		code:    code,
		showMsg: showMsg,
		error:   errors.New(msg),
	}
	return e
}

func (e *err) GetCode() int {
	return e.code
}

func (e *err) GetShowMsg() string {
	return e.showMsg
}
