package main

import (
	"fmt"
	"os"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

const (
	version = "0.0.1"
	name    = "chair"
	msg     = name + " is initial setup tool for isucon"
)

var perm os.FileMode = 0755

type Chair struct {
	Options
}

func main() {
	os.Exit((&Chair{}).Run())
}

// Run command line
func (c *Chair) Run() int {
	if e := c.run(); e != nil {
		exitCode, err := UnwrapErrors(e)
		if c.StackTrace {
			fmt.Fprintf(os.Stderr, "Error:\n  %+v\n", e)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		}
		return exitCode
	}
	return 0
}

func (c *Chair) run() error {
	_, err := c.prepare()
	if err != nil {
		return err
	}
	// Options
	if c.Show != "" {
		return c.runShow()
	}
	if c.Init {
		return runInit()
	}

	os.Stderr.Write(c.usage())
	return nil
}

func (c *Chair) prepare() ([]string, error) {
	args, err := parseOptions(&c.Options, os.Args[1:])
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse command line args")
	}
	return args, nil
}

func parseOptions(opts *Options, argv []string) ([]string, error) {
	o, err := opts.parse(argv)
	if err != nil {
		return nil, exit.MakeDataErr(err)
	}

	switch opts.Help {
	case "alp":
		return nil, errors.New("Try: sudo cat /var/log/nginx/access.log | alp -r")
	case "pt-query-digest":
		return nil, errors.New("Try: sudo pt-query-digest /var/log/mysql/mysqld-slow.log")
	}

	if opts.Version {
		return nil, exit.MakeUsage(errors.New(msg))
	}
	return o, nil
}
