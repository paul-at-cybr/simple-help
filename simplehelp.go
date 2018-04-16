package simplehelp

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

// SimpleHelp - object containing configuration options for the help function
type SimpleHelp struct {
	ProgramTitle       string // Title of your CLI program goes here
	ProgramDescription string // Program description
	Indentation        int    // How many characters to indent the flag descriptions
	helpFormatString   string
}

// Help - Gets available flags, combines with the config options,
// and outputs a help page to stdout.
// Usage: flag.CommandLine.Usage = help.Help
// Can also be called directly: help.Help()
func (h *SimpleHelp) Help() {

	h.helpFormatString = h.makeHelpFormatStr()

	grebo := color.New(color.FgGreen, color.Bold)
	grebo.Printf("\n%s\n", h.ProgramTitle)
	fmt.Println("  " + h.ProgramDescription)

	grebo.Print("\nFlags:\n")
	// print usage for all flags
	flag.CommandLine.VisitAll(func(fl *flag.Flag) {
		fmt.Printf("  --%s%s (Default: %s)\n", h.flagIndentation(fl.Name), fl.Usage, fl.DefValue)
	})

	fmt.Print("\n")
}

// creates a formatting string that can be used for space-padding via Sprintf
func (h *SimpleHelp) makeHelpFormatStr() string {
	return fmt.Sprintf("%%-%ds", h.Indentation)
}

// uses the formatting string to apply space padding
func (h *SimpleHelp) flagIndentation(flagName string) string {
	return fmt.Sprintf(h.helpFormatString, flagName)
}

// Hint: Prints a small hint about the --help parameter
func Hint() {
	fmt.Println("Use --help for more information")
}
