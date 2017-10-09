package main

import (
	"fmt"
	"os"

	"github.com/Code-Hex/chair/info"

	"github.com/Code-Hex/chair/dump"

	"github.com/Code-Hex/chair/tldr"

	"github.com/Code-Hex/chair/show"

	"github.com/Code-Hex/chair/setup"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
	name    = "chair"
	msg     = name + " is initial setup tool for isucon"
)

type Chair struct {
	StackTrace bool
	command    *cobra.Command
}

func New() *Chair {
	chair := new(Chair)
	chair.command = &cobra.Command{
		Use:           name,
		Short:         msg,
		Long:          msg,
		RunE:          chair.run,
		SilenceErrors: true,
	}

	chair.command.AddCommand(
		setup.CommandNew(),
		show.CommandNew(),
		tldr.CommandNew(),
		dump.CommandNew(),
		info.CommandNew(),
	)

	chair.command.Flags().Bool("version", true, "show version")

	return chair
}

func main() {
	os.Exit(New().Run())
}

// Run command line
func (c *Chair) Run() int {
	if e := c.command.Execute(); e != nil {
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

func (c *Chair) run(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
