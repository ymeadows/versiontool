package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SeeSpotRun/coerce"
	docopt "github.com/docopt/docopt-go"
	"github.com/samsalisbury/semv"
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

	debug("subcommand: %#v", sc)
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

func parseVersion(prefix, version string, strict bool) (semv.Version, error) {
	if version[0:len(prefix)] == prefix {
		version = version[len(prefix):len(version)]
	} else if strict {
		return semv.Version{}, fmt.Errorf("Version string: %q does not start with prefix %q", version, prefix)
	}

	v, err := semv.Parse(version)
	debug("version parsed: %#v", v)
	return v, err
}
