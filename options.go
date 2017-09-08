package main

import (
	"fmt"
)

type options struct {
	cmd       string
	args      []string
	url       string
	debug     bool
	mesosPort int
}

const docstring = `Issue commands against a Singularity
Usage: versiontool [options] <cmd> [<args>...]

Options:
  --url=<string>      The URL of the singularity to contact
  --debug, -d         Enable extra output
  --mesos-port=<int>  Mesos agent port [default: 5051]
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
