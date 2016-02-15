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

func (cmd createBaremetalsCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd createBaremetalsCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd createBaremetalsCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd createBaremetalsCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
