package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type slPackagesCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewSlPackagesCommand(options Options) slPackagesCommand {
	consoleUi := common.NewConsoleUi()

	return slPackagesCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd slPackagesCommand) Name() string {
	return "sl-packages"
}

func (cmd slPackagesCommand) Description() string {
	return "List all Softlayer packages"
}

func (cmd slPackagesCommand) Usage() string {
	return "bmp sl-packages"
}

func (cmd slPackagesCommand) Options() Options {
	return cmd.options
}

func (cmd slPackagesCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slPackagesCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
