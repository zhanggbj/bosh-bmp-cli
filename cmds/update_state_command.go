package cmds

import (
	"fmt"
)

type updateStateCommand struct {
	options Options
}

func NewUpdateStateCommand(options Options) updateStateCommand {
	return updateStateCommand{
		options: options,
	}
}

func (cmd updateStateCommand) Name() string {
	return "update-state"
}

func (cmd updateStateCommand) Description() string {
	return `Update the server state (\"bm.state.new\", \"bm.state.using\", \"bm.state.loading\", \"bm.state.failed\", \"bm.state.deleted\")`
}

func (cmd updateStateCommand) Usage() string {
	return "bmp update-state <bm.state.new>"
}

func (cmd updateStateCommand) Options() Options {
	return cmd.options
}

func (cmd updateStateCommand) Validate() (bool, error) {
	fmt.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd updateStateCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s command: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
