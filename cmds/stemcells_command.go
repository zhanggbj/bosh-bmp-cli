package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type stemcellsCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewStemcellsCommand(options Options) stemcellsCommand {
	consoleUi := common.NewConsoleUi()

	return stemcellsCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
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
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd stemcellsCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
