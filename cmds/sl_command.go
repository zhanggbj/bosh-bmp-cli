package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type slCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewSlCommand(options Options) slCommand {
	consoleUi := common.NewConsoleUi()

	return slCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd slCommand) Name() string {
	return "sl"
}

func (cmd slCommand) Description() string {
	return "List all Softlayer packages or package options"
}

func (cmd slCommand) Usage() string {
	return "bmp sl --packages | --package-options"
}

func (cmd slCommand) Options() Options {
	return cmd.options
}

func (cmd slCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
