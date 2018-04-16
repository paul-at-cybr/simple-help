package simplehelp

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

type SimpleHelp struct {
	ProgramTitle       string
	ProgramDescription string
	Indentation        int
	helpFormatString   string
}

func (h *SimpleHelp) Help() {

	h.helpFormatString = h.makeHelpFormatStr()

	grebo := color.New(color.FgGreen, color.Bold)
	grebo.Printf("\n%s\n", h.ProgramTitle)
	fmt.Println("  " + h.ProgramDescription)

	grebo.Print("\nFlags:\n")
	// print usage for all flags
	flag.CommandLine.VisitAll(func(fl *flag.Flag) {
		fmt.Printf("  --%s%s\n", h.flagIndentation(fl.Name), fl.Usage)
	})

	fmt.Print("\n")
}

func (h *SimpleHelp) makeHelpFormatStr() string {
	return fmt.Sprintf("%%-%ds", h.Indentation)
}

func (h *SimpleHelp) flagIndentation(flagName string) string {
	return fmt.Sprintf(h.helpFormatString, flagName)
}
