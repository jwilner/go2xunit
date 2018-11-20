// Command line parsing
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

var args struct {
	failRace    bool
	format      string
	input       string
	noFail      bool
	output      string
	suitePrefix string
	version     bool
}

func templateNames() string {
	names := make([]string, 0, len(internalTemplates))
	for name := range internalTemplates {
		names = append(names, name)
	}

	sort.Strings(names)
	return strings.Join(names, ", ")
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\nOptions:\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	formatHelp := fmt.Sprintf("output format: %s", templateNames())

	flag.BoolVar(&args.failRace, "fail-race", false, "mark test as failing if it exposes a data race")
	flag.BoolVar(&args.noFail, "no-fail", false, "don't fail if tests failed")
	flag.BoolVar(&args.version, "version", false, "print version and exit")
	flag.StringVar(&args.format, "format", "junit", formatHelp)
	flag.StringVar(&args.input, "input", "", "input file")
	flag.StringVar(&args.output, "output", "", "output file")
	flag.StringVar(&args.suitePrefix, "suite-prefix", "", "prefix to include before all suite names")
}

// parseArgs parses and validates command line arguments
func parseArgs() error {
	flag.Parse()

	if flag.NArg() > 0 {
		return fmt.Errorf("%s takes no arguments", os.Args[0])
	}

	return nil
}
