package chair

import (
	"fmt"
	"os"

	"github.com/Code-Hex/chair/internal/dump"
	"github.com/Code-Hex/chair/internal/info"
	"github.com/Code-Hex/chair/internal/setup"
	"github.com/Code-Hex/chair/internal/show"
	"github.com/Code-Hex/chair/internal/tldr"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
	name    = "chair"
	msg     = name + " is initial setup tool for isucon"
)

type Chair struct {
	StackTrace bool
	Command    *cobra.Command
}

func New() *Chair {
	chair := new(Chair)
	chair.Command = NewCommand()
	return chair
}

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:           name,
		Short:         msg,
		Long:          msg,
		RunE:          run(),
		SilenceErrors: true,
	}

	c.AddCommand(
		setup.CommandNew(),
		show.CommandNew(),
		tldr.CommandNew(),
		dump.CommandNew(),
		info.CommandNew(),
	)

	c.Flags().Bool("version", true, "show version")
	return c
}

// Run command line
func (c *Chair) Run() int {
	if e := c.Command.Execute(); e != nil {
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

func run() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	}
}
