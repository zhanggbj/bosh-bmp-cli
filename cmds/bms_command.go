package cmds

import (
	"fmt"
)

type bmsCommand struct {
	args    []string
	options Options
}

func NewBmsCommand(options Options) bmsCommand {
	return bmsCommand{
		options: options,
	}
}

func (cmd bmsCommand) Name() string {
	return "bms"
}

func (cmd bmsCommand) Description() string {
	return "List all bare metals"
}

func (cmd bmsCommand) Usage() string {
	return "bmp bms"
}

func (cmd bmsCommand) Options() Options {
	return cmd.options
}

func (cmd bmsCommand) Println(args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Println(args...)
	}

	return 0, nil
}

func (cmd bmsCommand) Printf(msg string, args ...interface{}) (int, error) {
	if cmd.options.Verbose {
		return fmt.Printf(msg, args...)
	}

	return 0, nil
}

func (cmd bmsCommand) Validate() (bool, error) {
	cmd.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd bmsCommand) Execute(args []string) (int, error) {
	cmd.Printf("Executing %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
