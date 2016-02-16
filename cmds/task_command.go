package cmds

import (
	common "github.com/maximilien/bosh-bmp-cli/common"
)

type taskCommand struct {
	args    []string
	options Options

	ui      common.UI
	printer common.Printer
}

func NewTaskCommand(options Options) taskCommand {
	consoleUi := common.NewConsoleUi()

	return taskCommand{
		options: options,

		ui:      consoleUi,
		printer: common.NewDefaultPrinter(consoleUi, options.Verbose),
	}
}

func (cmd taskCommand) Name() string {
	return "task"
}

func (cmd taskCommand) Description() string {
	return `Show the output of the task: \"option --debug, Get the debug info of the task\"`
}

func (cmd taskCommand) Usage() string {
	return "bmp task <task-id>"
}

func (cmd taskCommand) Options() Options {
	return cmd.options
}

func (cmd taskCommand) Validate() (bool, error) {
	cmd.printer.Printf("Validating %s command: args: %#v, options: %#v", cmd.Name(), cmd.args, cmd.options)
	return true, nil
}

func (cmd taskCommand) Execute(args []string) (int, error) {
	cmd.printer.Printf("Executing %s comamnd: args: %#v, options: %#v", cmd.Name(), args, cmd.options)
	return 0, nil
}
