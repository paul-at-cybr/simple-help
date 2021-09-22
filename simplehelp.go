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
	helpSections       []helpSection
	helpFormatString   string
}

type helpSection struct {
	name        string
	description string
}

// Help - Gets available flags, combines with the config options,
// and outputs a help page to stdout.
// Usage: flag.CommandLine.Usage = help.Help
// Can also be called directly: help.Help()
func (h *SimpleHelp) Help() {

	h.helpFormatString = h.makeHelpFormatStr()

	grebo := color.New(color.FgBlue, color.Bold)
	printSection(h.ProgramTitle, h.ProgramDescription, grebo)

	for _, section := range h.helpSections {
		printSection(section.name, section.description, grebo)
	}

	grebo.Print("\nFlags:\n")
	// print usage for all flags
	flag.CommandLine.VisitAll(func(fl *flag.Flag) {
		var defaultString string
		if fl.DefValue != "" {
			defaultString = fmt.Sprintf("(Default: %s)", fl.DefValue)
		}
		fmt.Printf("  --%s%s %s\n", h.flagIndentation(fl.Name), fl.Usage, defaultString)
	})

	fmt.Print("\n")
}

// AddSection - Adds a help section to the help output.
func (h *SimpleHelp) AddSection(name string, description string) {
	h.helpSections = append(h.helpSections, helpSection{name: name, description: description})
}

// Prints a section (colored title, plain content)
func printSection(title string, content string, c *color.Color) {
	c.Printf("\n%s\n", title)
	fmt.Println("  " + content)
}

// creates a formatting string that can be used for space-padding via Sprintf
func (h *SimpleHelp) makeHelpFormatStr() string {
	return fmt.Sprintf("%%-%ds", h.Indentation)
}

// uses the formatting string to apply space padding
func (h *SimpleHelp) flagIndentation(flagName string) string {
	return fmt.Sprintf(h.helpFormatString, flagName)
}

// Hint - Prints a small hint about the --help parameter
func Hint() {
	fmt.Println("Use --help for more information")
}
