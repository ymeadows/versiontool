package main

import "log"

var subCommands = map[string]subCommand{
	"help":      &helpCmd{},
	"increment": &incrementCmd{},
	"decrement": &decrementCmd{},
	"cut":       &cutCommand{},
	"sort":      &sortCmd{},
}

func main() {
	log.SetFlags(0)
	parsed := parseOpts(nil)
	opts := options{}
	parsed.fill(&opts)

	cmd := chooseSubcommand(opts.cmd)

	run(opts.cmd, cmd, opts.args, parsed)
}

func chooseSubcommand(name string) subCommand {
	if cmd, has := subCommands[name]; has {
		return cmd
	}
	log.Fatalf("Subcommand not recognized: %s", name)
	return nil
}
