package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type bmsCommand struct {
	args    []string
	options Options

	ui      common.UI
	printer common.Printer
}

func NewBmsCommand(options Options) bmsCommand {
	consoleUi := common.NewConsoleUi()

	return bmsCommand{
		options: options,
		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
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

func (cmd bmsCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd bmsCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return 0, nil
}
