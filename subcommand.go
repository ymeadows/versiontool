package main

import (
	"log"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
)

type subCommand interface {
	description() string
	docs() string
	action()
}

func run(name string, sc subCommand, args []string, toplevel map[string]interface{}) {
	argv := []string{name}
	argv = append(argv, args...)
	subParseArgv(sc, argv, toplevel)

	sc.action()
}

func fullDocs(docs string) string {
	return docs + `
For common options, see 'versiontool help'
`
}

func subParseArgv(sc subCommand, argv []string, toplevel map[string]interface{}) {
	parsed, err := docopt.Parse(fullDocs(sc.docs()), argv, true, "", false)

	if err != nil {
		log.Fatal(err)
	}

	err = coerce.Struct(sc, toplevel, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
	err = coerce.Struct(sc, parsed, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
}
