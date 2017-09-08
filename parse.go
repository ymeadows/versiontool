package main

import (
	"log"

	docopt "github.com/docopt/docopt-go"
)

type parsed map[string]interface{}

func parseOpts(argv []string) parsed {
	p, err := docopt.Parse(docstring, argv, true, "versiontool v0.0.1", false)

	if err != nil {
		log.Fatal(err)
	}

	return p
}

func (p parsed) fill(opts interface{}) {
	err = coerce.Struct(opts, p, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
}
