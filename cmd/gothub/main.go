package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/elimity-com/backend-intern-exercise/internal"
	"github.com/google/go-github/v33/github"
)

var args = os.Args

var name = makeName()

func log(message string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, message)
}

func main() {
	err, usage := run()
	if err != "" {
		log(err)
		if usage {
			message := fmt.Sprintf("run '%s help' for usage information", name)
			log(message)
		}
		os.Exit(1)
	}
}

func makeName() string {
	path := args[0]
	return filepath.Base(path)
}

func parseInterval() (time.Duration, string) {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	var interval time.Duration
	set.DurationVar(&interval, "interval", 10*time.Second, "")
	set.SetOutput(ioutil.Discard)
	args := args[2:]
	if err := set.Parse(args); err != nil {
		return 0, "got invalid flags"
	}
	if interval <= 0 {
		return 0, "got invalid interval"
	}
	return interval, ""
}

func run() (string, bool) {
	if nbArgs := len(args); nbArgs < 2 {
		return "missing command", true
	}
	switch args[1] {
	case "help":
		const usage = `
Simple CLI for tracking public GitHub repositories.

Usage:
  %[1]s help
  %[1]s track [-interval=<interval>]

Commands:
  help  Show usage information
  track Track public GitHub repositories

Options:
  -interval=<interval> Repository update interval, greater than zero [default: 10s]
`
		fmt.Fprintf(os.Stdout, usage, name)
		return "", false

	case "track":
		client := github.NewClient(nil)
		interval, err := parseInterval()
		if err != "" {
			err := fmt.Sprintf("failed parsing interval: %s", err)
			return err, true
		}
		if err := internal.Track(client.Search.Repositories, interval); err != nil {
			err := fmt.Sprintf("failed tracking: %v", err)
			return err, false
		}
		return "", false

	default:
		return "got invalid command", true
	}
}
