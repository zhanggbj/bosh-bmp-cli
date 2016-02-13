package cmds

import "fmt"

type updateStateCommand struct {
	args    []string
	options Options
}

func NewUpdateStateCommand(args []string, options Options) updateStateCommand {
	return updateStateCommand{
		args:    args,
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

func (cmd updateStateCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
