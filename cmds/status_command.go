package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type statusCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewStatusCommand(options Options) statusCommand {
	consoleUi := common.NewConsoleUi()

	return statusCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd statusCommand) Name() string {
	return "status"
}

func (cmd statusCommand) Description() string {
	return "show bmp status"
}

func (cmd statusCommand) Usage() string {
	return "bmp status"
}

func (cmd statusCommand) Options() Options {
	return cmd.options
}

func (cmd statusCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd statusCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
