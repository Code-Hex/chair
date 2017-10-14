package tldr

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

const indent = "    "

type TLDR interface {
	Print()
	Name() string
	Match(string) bool
}

var cmds = []TLDR{
	alpNew(),
	ptQueryDigestNew(),
}

func CommandNew() *cobra.Command {
	return &cobra.Command{
		Use:       "tldr",
		Short:     "tldr some command",
		Long:      "tldr some command",
		Example:   example(),
		RunE:      run,
		ValidArgs: validArgs(),
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		for _, c := range cmds {
			if c.Match(args[0]) {
				c.Print()
				return nil
			}
		}
	}
	return cmd.Usage()
}

func example() string {
	base := indent + "chair tldr "
	var buf bytes.Buffer
	for _, c := range cmds {
		fmt.Fprintf(&buf, base+c.Name()+"\n")
	}
	r := buf.String()
	return r[:len(r)-1] // trim newline
}

func validArgs() []string {
	args := make([]string, 0, len(cmds))
	for _, c := range cmds {
		args = append(args, c.Name())
	}
	return args
}
