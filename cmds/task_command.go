package cmds

import (
	"fmt"
)

type taskCommand struct {
	args    []string
	options Options
}

func NewTaskCommand(options Options) taskCommand {
	return taskCommand{
		options: options,
	}
}

func (cmd taskCommand) Name() string {
	return "task"
}

func (cmd taskCommand) Description() string {
	return `Show the output of the task: \"option --debug, Get the debug info of the task\"`
}

func (cmd taskCommand) Usage() string {
	return "bmp task <task-id>"
}

func (cmd taskCommand) Options() Options {
	return cmd.options
}

func (cmd taskCommand) Validate() (bool, error) {
	fmt.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd taskCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
