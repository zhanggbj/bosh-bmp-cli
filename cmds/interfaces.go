package cmds

type Options struct {
	Help    string `short:"h" long:"help" description:"Show help information"`
	Verbose bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
}

type Printer interface {
	Println(args ...interface{}) (int, error)
	Printf(msg string, args ...interface{}) (int, error)
}

type Command interface {
	Printer

	Name() string
	Description() string
	Usage() string
	Options() Options

	Validate() (bool, error)
	Execute(args []string) (int, error)
}
