package cmds

import (
	"fmt"
)

type slPackageOptionsCommand struct {
	options Options
}

func NewSlPackageOptionsCommand(options Options) slPackageOptionsCommand {
	return slPackageOptionsCommand{
		options: options,
	}
}

func (cmd slPackageOptionsCommand) Name() string {
	return "sl-package-options"
}

func (cmd slPackageOptionsCommand) Description() string {
	return "List all options of Softlayer package"
}

func (cmd slPackageOptionsCommand) Usage() string {
	return "bmp sl-package-options"
}

func (cmd slPackageOptionsCommand) Options() Options {
	return cmd.options
}

func (cmd slPackageOptionsCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd slPackageOptionsCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd slPackageOptionsCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slPackageOptionsCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
