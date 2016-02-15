package cmds

import (
	"fmt"
)

type loginCommand struct {
	options Options
}

func NewLoginCommand(options Options) loginCommand {
	return loginCommand{
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

func (cmd loginCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd loginCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd loginCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd loginCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
