package main

import (
	"fmt"
	"io"
	"log"

	"github.com/samsalisbury/semv"
)

type cutCommand struct {
	major, minor, patch bool
	version             string
	prefix              string
	strict              bool
}

func (cut *cutCommand) description() string {
	return "increments a version"
}

func (cut *cutCommand) docs() string {
	return cut.description() + `
Usage: cut [options] <version>

Options:
	--major, -M  Cut after major version
	--minor, -m  Cut after minor version
	--patch, -p  Cut after patch version

Example:
> versiontool cut -M 1.2.3-rc1
1

> versiontool cut -m 1.2.3-rc1
1.2

> versiontool cut -p 1.2.3
1.2.3

`
}

func (cut *cutCommand) action(out, _ io.Writer) {
	version, err := parseVersion(cut.prefix, cut.version, cut.strict)
	if err != nil {
		log.Fatal(err)
	}

	if cut.major {
		fmt.Fprintf(out, "%s%s\n", cut.prefix, version.Format(semv.Major))
		return
	}

	if cut.minor {
		fmt.Fprintf(out, "%s%s\n", cut.prefix, version.Format(semv.MajorMinor))
		return
	}

	fmt.Fprintf(out, "%s%s\n", cut.prefix, version.Format(semv.MajorMinorPatch))
}
