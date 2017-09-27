package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	sorting "sort"

	"github.com/samsalisbury/semv"
)

type (
	sortCmd struct{}

	versionList []semv.Version
)

func (sort sortCmd) description() string {
	return "sorts a list of version from stdin"
}

func (sort sortCmd) docs() string {
	return sort.description() + `
Usage: sort
	`
}

func (vl versionList) Len() int {
	return len(vl)
}

func (vl versionList) Less(i int, j int) bool {
	return vl[i].Less(vl[j])
}

func (vl versionList) Swap(i int, j int) {
	vl[i], vl[j] = vl[j], vl[i]
}

func (sort sortCmd) action(out, err io.Writer) {
	prefixes := regexp.MustCompile(`^[^\d]*`)
	lines := bufio.NewScanner(os.Stdin)
	versions := versionList{}
	for lines.Scan() {
		versiontext := prefixes.ReplaceAllLiteralString(lines.Text(), "")
		version, err := semv.Parse(versiontext)
		if err != nil {
			// implicit filter of badly formed versions
			continue
		}
		versions = append(versions, version)
	}

	sorting.Stable(versions)

	for _, v := range versions {
		fmt.Fprintln(out, v)
	}
}
