package cmds

import (
	"errors"
	"fmt"
)

type stemcellsCommand struct {
	args    []string
	options Options
}

func NewStemcellsCommand(args []string, options Options) stemcellsCommand {
	return stemcellsCommand{
		args:    args,
		options: options,
	}
}

func (cmd stemcellsCommand) Name() string {
	return "stemcells"
}

func (cmd stemcellsCommand) Description() string {
	return "List all stemcells"
}

func (cmd stemcellsCommand) Usage() string {
	return "bmp stemcells"
}

func (cmd stemcellsCommand) Options() Options {
	return cmd.options
}

func (cmd stemcellsCommand) Validate() (bool, error) {
	return false, errors.New("Implement me!")
}

func (cmd stemcellsCommand) Execute(args []string) (int, error) {
	fmt.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
