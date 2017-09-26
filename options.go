package main

import (
	"fmt"
)

type options struct {
	cmd    string
	args   []string
	prefix string
	strict bool
}

const docstring = `Process and manipulate semantic version strings
Usage: versiontool [options] <cmd> [<args>...]

Options:
  --prefix=<string>   A prefix to use with versions. (e.g "v" for v1.2.3)
	--strict, -S        Absolutely require the presence of the prefix
`

func docStr() string {
	str := docstring + "\nMost common subcommands:"
	for _, k := range []string{"help", "get-logs"} {
		str = str + fmt.Sprintf("\n  %s: %s",
			k, subCommands[k].description)
	}
	str = str + "\n"
	return str
}
