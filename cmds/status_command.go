package cmds

import (
	"errors"
	"fmt"
)

type statusCommand struct {
	args    []string
	options Options
}

func NewStatusCommand(args []string, options Options) statusCommand {
	return statusCommand{
		args:    args,
		options: options,
	}
}

func (cmd statusCommand) Name() string {
	return "status"
}

func (cmd statusCommand) Description() string {
	return "show bmp status"
}

func (cmd statusCommand) Usage() string {
	return "bmp status"
}

func (cmd statusCommand) Options() Options {
	return cmd.options
}

func (cmd statusCommand) Validate() (bool, error) {
	return false, errors.New("Implement me!")
}

func (cmd statusCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
