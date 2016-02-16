package fakes

import (
	"fmt"

	common "github.com/maximilien/bosh-bmp-cli/common"
)

type FakeUi struct {
	Output string

	Msg  string
	Args []interface{}

	PrintfRc  int
	PrintfErr error

	PrintlnRc  int
	PrintlnErr error
}

func NewFakeUi() common.UI {
	return FakeUi{
		Output: "",
		Msg:    "",
		Args:   make([]interface{}, 0),
	}
}

func (fakeUi FakeUi) Printf(msg string, args ...interface{}) (int, error) {
	fakeUi.Msg = msg
	fakeUi.Args = args

	fakeUi.Output = fmt.Sprintf(msg, args...)

	return fakeUi.PrintfRc, fakeUi.PrintfErr
}

func (fakeUi FakeUi) Println(args ...interface{}) (int, error) {
	fakeUi.Args = args

	fakeUi.Output = fmt.Sprintf("%s\n", args...)

	return fakeUi.PrintlnRc, fakeUi.PrintlnErr
}
