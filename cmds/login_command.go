package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type loginCommand struct {
	options Options

	ui      common.UI
	printer common.Printer
}

func NewLoginCommand(options Options) loginCommand {
	consoleUi := common.NewConsoleUi()

	return loginCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd loginCommand) Name() string {
	return "login"
}

func (cmd loginCommand) Description() string {
	return "Login to the Bare Metal Provision Server"
}

func (cmd loginCommand) Usage() string {
	return "bmp login"
}

func (cmd loginCommand) Options() Options {
	return cmd.options
}

func (cmd loginCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: options: %#v", cmd.Name(), cmd.options)
	return true, nil
}

func (cmd loginCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
