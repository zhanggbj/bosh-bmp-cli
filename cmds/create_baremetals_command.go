package cmds

import (
	"fmt"
)

type createBaremetalsCommand struct {
	options Options
}

func NewCreateBaremetalsCommand(options Options) createBaremetalsCommand {
	return createBaremetalsCommand{
		options: options,
	}
}

func (cmd createBaremetalsCommand) Name() string {
	return "create-baremetals"
}

func (cmd createBaremetalsCommand) Description() string {
	return `Create the missed baremetals: \"option --dryrun, only verify the orders\"`
}

func (cmd createBaremetalsCommand) Usage() string {
	return "bmp create-baremetals [--dryrun]"
}

func (cmd createBaremetalsCommand) Options() Options {
	return cmd.options
}

func (cmd createBaremetalsCommand) Validate() (bool, error) {
	fmt.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd createBaremetalsCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
