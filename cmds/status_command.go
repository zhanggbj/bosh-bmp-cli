package cmds

import (
	"fmt"
)

type statusCommand struct {
	options Options
}

func NewStatusCommand(options Options) statusCommand {
	return statusCommand{
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
	fmt.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd statusCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
