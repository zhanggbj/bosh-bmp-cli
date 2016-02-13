package cmds

type Options struct {
	Help    string `short:"h" long:"help" description:"Show help information"`
	Verbose bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
}

type Command interface {
	Name() string
	Description() string
	Usage() string
	Options() Options
	Validate() (bool, error)
	Execute(args []string) (int, error)
}
