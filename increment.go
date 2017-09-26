package main

import "github.com/samsalisbury/semv"

type incrementCmd struct{}

func (inc *incrementCmd) description() string {
	return "increments a version"
}

func (inc *incrementCmd) docs() string {
	return inc.description() + `
  Usage: versiontool increment [options] <version>

  Options:
    --no-reset, -R            Do not reset smaller increments
    --major=<int>, -M<int>    Increment major version by <int>
    --minor=<int>, -m<int>    Increment minor version by <int>
    --patch=<int>, -p<int>    Increment patch version by <int>

  Resetting will clear or set to zero parts of the version less significant
  than the parts being incremented. This is almost always what.

  With no arguments provided, the default is -p1 - that is to bump the patch
  version by one.

  Example:
    > versiontool increment -M 1 1.2.3-rc1
    2.0.0

    > versiontool increment -M 1 -m 1 1.2.3-rc1
    2.1.0

    > versiontool increment --no-reset -M 1 1.2.3-rc1
    2.2.3-rc1

    > versiontool increment --no-reset -M 1 -m 1 1.2.3-rc1
    2.3.3-rc1
  `
}

func (inc *incrementCmd) action() {
	version := semv.MustParse(inc.version)

	if inc.major == 0 && inc.minor == 0 && inc.patch == 0 {
		inc.patch = 1
	}

	if inc.major > 0 {
		version.Major += inc.major
		if !inc.noReset {
			version.Minor = 0
			version.Patch = 0
			version.Meta = ""
			version.Pre = ""
		}
	}

	if inc.minor > 0 {
		version.Minor += inc.minor
		if !inc.noReset {
			version.Patch = 0
			version.Meta = ""
			version.Pre = ""
		}
	}

	if inc.patch > 0 {
		version.Patch += inc.patch
		if !inc.noReset {
			version.Meta = ""
			version.Pre = ""
		}
	}

	fmt.Println(version)
}
