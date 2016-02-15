package cmds

import (
	"fmt"
)

type targetCommand struct {
	args    []string
	options Options
}

func NewTargetCommand(options Options) targetCommand {
	return targetCommand{
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

func (cmd targetCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd targetCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd targetCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd targetCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
