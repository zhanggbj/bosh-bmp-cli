package cmds

import "fmt"

type createBaremetalsCommand struct {
	args    []string
	options Options
}

func NewCreateBaremetalsCommand(args []string, options Options) createBaremetalsCommand {
	return createBaremetalsCommand{
		args:    args,
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

func (cmd createBaremetalsCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
