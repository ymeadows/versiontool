package main

import "io"

type helpCmd struct {
	cmd string
}

func (help *helpCmd) description() string {
	return "prints help"
}

func (help *helpCmd) docs() string {
	return help.description() + `
	Usage: versiontool help [<cmd>]
	`
}

func (help *helpCmd) action(out io.Writer, err io.Writer) {
	switch help.cmd {
	default:
		subParseArgv(chooseSubcommand(help.cmd), []string{"--help"}, map[string]interface{}{})
	case "":
		parseOpts([]string{"--help"})
	}
}
