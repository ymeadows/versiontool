package main

import "log"

var subCommands = map[string]subCommand{
	"help":      &helpCmd{},
	"increment": &incrementCmd{},
	"decrement": &decrementCmd{},
	"cut":       &cutCommand{},
	"sort":      &sortCmd{},
	"highest":   &highestCmd{},
}

var debug = func(fmt string, args ...interface{}) {
}

func main() {
	log.SetFlags(0)
	parsed := parseOpts(nil)
	opts := options{}
	parsed.fill(&opts)

	cmd := chooseSubcommand(opts.cmd)

	if opts.debug {
		debug = func(fmt string, args ...interface{}) {
			log.Printf(fmt, args)
		}
	}

	debug("Top level args: %#v", opts)
	run(opts.cmd, cmd, opts.args, parsed)
}

func chooseSubcommand(name string) subCommand {
	if cmd, has := subCommands[name]; has {
		return cmd
	}
	log.Fatalf("Subcommand not recognized: %s", name)
	return nil
}
