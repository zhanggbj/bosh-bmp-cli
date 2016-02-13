package cmds

import "fmt"

type targetCommand struct {
	args    []string
	options Options
}

func NewTargetCommand(args []string, options Options) targetCommand {
	return targetCommand{
		args:    args,
		options: options,
	}
}

func (cmd targetCommand) Name() string {
	return "target"
}

func (cmd targetCommand) Description() string {
	return "Set the URL of Bare Metal Provision Server"
}

func (cmd targetCommand) Usage() string {
	return "bmp target http://url.to.bmp.server"
}

func (cmd targetCommand) Options() Options {
	return cmd.options
}

func (cmd targetCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
