package common

type defaultPrinter struct {
	Ui      UI
	Verbose bool
}

func NewDefaultPrinter(ui UI, verbose bool) Printer {
	return defaultPrinter{
		Ui:      ui,
		Verbose: verbose,
	}
}

func (p defaultPrinter) Printf(msg string, args ...interface{}) (int, error) {
	if p.Verbose {
		return p.Ui.Printf(msg, args)
	}

	return 0, nil
}

func (p defaultPrinter) Println(args ...interface{}) (int, error) {
	if p.Verbose {
		return p.Ui.Println(args)
	}

	return 0, nil
}
