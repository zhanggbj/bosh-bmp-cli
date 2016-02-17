package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"runtime/debug"

	flags "github.com/jessevdk/go-flags"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var options cmds.Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	args, err := parser.ParseArgs(os.Args)
	if err != nil {
		handlePanic()
		os.Exit(1)
	}

	command, err := createCommand(args, options)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	validated, err := command.Validate()
	if err != nil {
		handlePanic()
		os.Exit(1)
	}

	if !validated {
		fmt.Println("Invalid options for command: ", err.Error())
		os.Exit(1)
	}

	rc, err := command.Execute(args)
	if err != nil {
		handlePanic()
		os.Exit(rc)
	}

	os.Exit(rc)
}

func createCommand(args []string, options cmds.Options) (cmds.Command, error) {
	if len(args) < 2 {
		return nil, errors.New("No bmp command specified")
	}

	cmd := createCommands(options)[args[1]]
	if cmd == nil {
		return nil, errors.New(fmt.Sprintf("Invalid command: %s", args[1]))
	}

	return cmd, nil
}

func createCommands(options cmds.Options) map[string]cmds.Command {
	return map[string]cmds.Command{
		"bms":               cmds.NewBmsCommand(options),
		"create-baremetals": cmds.NewCreateBaremetalsCommand(options),
		"login":             cmds.NewLoginCommand(options),
		"sl":                cmds.NewSlCommand(options),
		"status":            cmds.NewStatusCommand(options),
		"stemcells":         cmds.NewStemcellsCommand(options),
		"target":            cmds.NewTargetCommand(options),
		"task":              cmds.NewTaskCommand(options),
		"tasks":             cmds.NewTasksCommand(options),
		"update-state":      cmds.NewUpdateStateCommand(options),
	}
}

func handlePanic() {
	err := recover()

	if err != nil {
		switch err := err.(type) {
		case error:
			displayCrashDialog(err.Error())
		case string:
			displayCrashDialog(err)
		default:
			displayCrashDialog("An unexpected type of error")
		}
	}

	if err != nil {
		os.Exit(1)
	}
}

func displayCrashDialog(errorMessage string) {
	formattedString := `
Something completely unexpected happened. This is a bug in %s.
Please file this bug : https://github.com/maximilien/bosh-bmp-cli/issues
Tell us that you ran this command:

	%s

this error occurred:

	%s

and this stack trace:

%s
	`

	stackTrace := "\t" + strings.Replace(string(debug.Stack()), "\n", "\n\t", -1)
	println(fmt.Sprintf(formattedString, "bosh-bmp-cli", strings.Join(os.Args, " "), errorMessage, stackTrace))
}
