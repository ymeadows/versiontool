package main

import (
	"fmt"
	"io"

	"github.com/samsalisbury/semv"
)

type decrementCmd struct {
	major, minor, patch int
	reset               bool
	version             string
}

func (dec *decrementCmd) description() string {
	return "decrements a version"
}

func (dec *decrementCmd) docs() string {
	return dec.description() + `
Usage: decrement [options] <version>

Options:
	--reset, -r               Reset smaller decrements
	--major=<int>, -M<int>    Increment major version by <int>
	--minor=<int>, -m<int>    Increment minor version by <int>
	--patch=<int>, -p<int>    Increment patch version by <int>

Resetting will clear or set to zero parts of the version less significant
than the parts being decremented.

With no arguments provided, the default is -p1 - that is to bump down the patch
version by one.

Example:
	> versiontool decrement  -M 1 2.2.3-rc1
	1.2.3-rc1

	> versiontool decrement  -M 1 -m 1 2.2.3-rc1
	1.1.3-rc1

	> versiontool decrement --reset -M 1 2.2.3-rc1
	1.0.0

	> versiontool decrement --reset -M 1 -m 1 2.2.3-rc1
	1.1.0
`
}

func (dec *decrementCmd) action(out, err io.Writer) {
	version := semv.MustParse(dec.version)

	if dec.major == 0 && dec.minor == 0 && dec.patch == 0 {
		dec.patch = 1
	}

	if dec.major > 0 {
		version.Major -= dec.major
		if dec.reset {
			version.Minor = 0
			version.Patch = 0
			version.Meta = ""
			version.Pre = ""
		}
	}

	if dec.minor > 0 {
		version.Minor -= dec.minor
		if dec.reset {
			version.Patch = 0
			version.Meta = ""
			version.Pre = ""
		}
	}

	if dec.patch > 0 {
		version.Patch -= dec.patch
		if dec.reset {
			version.Meta = ""
			version.Pre = ""
		}
	}

	fmt.Fprintln(out, version)
}
