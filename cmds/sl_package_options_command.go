package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type slPackageOptionsCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewSlPackageOptionsCommand(options Options) slPackageOptionsCommand {
	consoleUi := common.NewConsoleUi()

	return slPackageOptionsCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd slPackageOptionsCommand) Name() string {
	return "sl-package-options"
}

func (cmd slPackageOptionsCommand) Description() string {
	return "List all options of Softlayer package"
}

func (cmd slPackageOptionsCommand) Usage() string {
	return "bmp sl-package-options"
}

func (cmd slPackageOptionsCommand) Options() Options {
	return cmd.options
}

func (cmd slPackageOptionsCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd slPackageOptionsCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
