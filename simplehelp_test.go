package simplehelp

import (
	"flag"
	"testing"
)

var count int64

func TestCreate(t *testing.T) {
	flag.Int64Var(&count, "count", 0, "Number of sheep to count")

	help := &SimpleHelp{
		ProgramTitle:       "Test program",
		ProgramDescription: "Test program description",
		Indentation:        30,
	}

	flag.CommandLine.Usage = help.Help

	flag.Parse()

	help.Help()
}
