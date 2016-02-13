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

func (cmd slPackagesCommand) Validate() (bool, error) {
	fmt.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slPackagesCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
