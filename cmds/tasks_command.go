package cmds

import (
	"fmt"
)

type tasksCommand struct {
	args    []string
	options Options
}

func NewTasksCommand(options Options) tasksCommand {
	return tasksCommand{
		options: options,
	}
}

func (cmd tasksCommand) Name() string {
	return "tasks"
}

func (cmd tasksCommand) Description() string {
	return `List tasks: \"option --latest number\", \"The number of latest tasks, default is 50\"`
}

func (cmd tasksCommand) Usage() string {
	return "bmp tasks"
}

func (cmd tasksCommand) Options() Options {
	return cmd.options
}

func (cmd tasksCommand) Validate() (bool, error) {
	fmt.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd tasksCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
