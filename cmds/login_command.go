package cmds

import "fmt"

type loginCommand struct {
	args    []string
	options Options
}

func NewLoginCommand(args []string, options Options) loginCommand {
	return loginCommand{
		args:    args,
		options: options,
	}
}

func (cmd loginCommand) Name() string {
	return "login"
}

func (cmd loginCommand) Description() string {
	return "Login to the Bare Metal Provision Server"
}

func (cmd loginCommand) Usage() string {
	return "bmp login"
}

func (cmd loginCommand) Options() Options {
	return cmd.options
}

func (cmd loginCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
