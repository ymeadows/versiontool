package main

import "log"

var subCommands = map[string]*subCommand{
	"help":      &helpCmd(),
	"increment": &incrementCmd(),
	"decrement": &decrementCmd(),
	"sort":      &sortCmd(),
}

func main() {
	log.SetFlags(0)
	parsed := parseOpts()
	opts := options{}
	parsed.fill(&opts)

	var cmd subCommand

	switch opts.cmd {
	default:
		log.Fatalf("Subcommand not recognized: %s", opts.cmd)
	case "help":
		cmd = &helpCmd{}
	case "increment":
		cmd = &incrementCmd{}
	case "decrement":
		cmd = &decrementCmd{}
	case "sort":
		cmd = &sortCmd{}
	}

	run(cmd, opts.cmd, cpts.args, parsed)
}
