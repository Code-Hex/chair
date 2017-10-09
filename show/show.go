package show

import (
	"fmt"
	"os"

	"github.com/Code-Hex/chair/show/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	indent             = "    "
	perm   os.FileMode = 0755
)

func CommandNew() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show some configuration content",
		Long:  "Show some configuration content",
		Example: fmt.Sprintf(
			"%s%s\n%s%s\n%s%s\n%s%s",
			indent,
			"chair show sql-default",
			indent,
			"chair show nginx-access-log",
			indent,
			"chair show sql-(default|slow-log|maybe-nice|cache|fix57-groupby)",
			indent,
			"chair show nginx-(access-log|event|outside-maybe-nice|static)",
		),
		RunE: run,
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		switch args[0] {
		// MySQL
		case "sql-default":
			fmt.Println(config.SQLDefault())
		case "sql-slow-log":
			fmt.Println(config.SQLSlowLog())
		case "sql-maybe-nice":
			dumpStr, err := config.SQLMaybeNice()
			if err != nil {
				return errors.Wrap(err, "Failed to get memory size")
			}
			fmt.Println(dumpStr)
		case "sql-cache":
			fmt.Println(config.SQLCache())
		case "sql-fix57-groupby":
			fmt.Println(config.SQLFix57GroupByProblem())
		// Nginx
		case "nginx-access-log":
			fmt.Println(config.NginxAccessLog())
		case "nginx-event":
			fmt.Println(config.NginxEvent())
		case "nginx-outside-maybe-nice":
			fmt.Println(config.NginxOutsideMaybeNice())
		case "nginx-static":
			fmt.Println(config.NginxStaticLocation())
		default:
			return cmd.Usage()
		}
		return nil
	}
	return cmd.Usage()
}
