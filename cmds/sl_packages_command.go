package cmds

import (
	"fmt"
)

type slPackagesCommand struct {
	options Options
}

func NewSlPackagesCommand(options Options) slPackagesCommand {
	return slPackagesCommand{
		options: options,
	}
}

func (cmd slPackagesCommand) Name() string {
	return "sl-packages"
}

func (cmd slPackagesCommand) Description() string {
	return "List all Softlayer packages"
}

func (cmd slPackagesCommand) Usage() string {
	return "bmp sl-packages"
}

func (cmd slPackagesCommand) Options() Options {
	return cmd.options
}

func (cmd slPackagesCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd slPackagesCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd slPackagesCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slPackagesCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
