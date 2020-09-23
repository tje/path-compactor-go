package main

import (
	"os"
	"strings"
	"regexp"
	"fmt"
)

func main() {
	var dir string
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			// @todo How should this be handled?
			return
		}
		dir = cwd
	}

	home, err := os.UserHomeDir()
	if err == nil {
		dir = strings.Replace(dir, home, "~", 1)
	}

	parts := strings.Split(dir, string(os.PathSeparator))
	rxSplit := regexp.MustCompile("[-_.]|[a-z](?:[A-Z])")
	for i, part := range parts {
		if len(part) <= 4 {
			continue
		}

		frags := rxSplit.Split(part, -1)
		if len(frags) == 1 {
			parts[i] = part[:3]
			continue
		}
		for fi, f := range frags {
			frags[fi] = f[:1]
		}
		parts[i] = strings.Join(frags, "")
	}

	fmt.Print(strings.Join(parts, string(os.PathSeparator)))
}
