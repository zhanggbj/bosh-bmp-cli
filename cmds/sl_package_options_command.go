package cmds

import "fmt"

type slPackageOptionsCommand struct {
	args    []string
	options Options
}

func NewSlPackageOptionsCommand(args []string, options Options) slPackageOptionsCommand {
	return slPackageOptionsCommand{
		args:    args,
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

func (cmd slPackageOptionsCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
