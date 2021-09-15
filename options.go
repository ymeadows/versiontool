package main

import (
	"fmt"
)

type options struct {
	cmd    string
	args   []string
	prefix string
	strict bool
	debug  bool
}

const docstring = `Process and manipulate semantic version strings
Usage: versiontool [options] <cmd> [<args>...]

Options:
  --prefix=<string>   A prefix to use with versions. (e.g "v" for v1.2.3)
  --strict, -S        Absolutely require the presence of the prefix
	--debug             Debug output
`

func docStr() string {
	str := docstring + "\nMost common subcommands:"
	for name, cmd := range subCommands {
		str = str + fmt.Sprintf("\n  %s: %s", name, cmd.description())
	}
	str = str + "\n"
	return str
}
