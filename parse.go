package main

import (
	"log"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
)

type parsed map[string]interface{}

func parseOpts(argv []string) parsed {
	p, err := docopt.Parse(docStr(), argv, true, "versiontool v0.0.1", true)

	if err != nil {
		log.Fatal(err)
	}

	return p
}

func (p parsed) fill(opts interface{}) {
	err := coerce.Struct(opts, p, "-%s", "--%s", "<%s>")
	if err != nil {
		log.Fatal(err)
	}
}
