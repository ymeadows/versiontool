package main

import (
	"io"
	"log"
	"os"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
)

type subCommand interface {
	description() string
	docs() string
	action(out, err io.Writer)
}

func run(name string, sc subCommand, args []string, toplevel map[string]interface{}) {
	if err := subParseArgv(sc, args, toplevel); err != nil {
		log.Fatal(err)
	}

	sc.action(os.Stdout, os.Stderr)
}

func fullDocs(docs string) string {
	return docs + `
For common options, see 'versiontool help'
`
}

func subParseArgv(sc subCommand, argv []string, toplevel map[string]interface{}) error {
	parsed, err := docopt.Parse(fullDocs(sc.docs()), argv, true, "", true)

	if err != nil {
		return err
	}

	err = coerce.Struct(sc, toplevel, "-%s", "--%s", "<%s>")
	if err != nil {
		return err
	}
	return coerce.Struct(sc, parsed, "-%s", "--%s", "<%s>")
}
