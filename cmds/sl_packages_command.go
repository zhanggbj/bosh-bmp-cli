package cmds

import "fmt"

type slPackagesCommand struct {
	args    []string
	options Options
}

func NewSlPackagesCommand(args []string, options Options) slPackagesCommand {
	return slPackagesCommand{
		args:    args,
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

func (cmd slPackagesCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
