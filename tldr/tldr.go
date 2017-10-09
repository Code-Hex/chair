package tldr

import (
	"fmt"

	"github.com/spf13/cobra"
)

const indent = "    "

func CommandNew() *cobra.Command {
	return &cobra.Command{
		Use:   "tldr",
		Short: "tldr some command",
		Long:  "tldr some command",
		Example: fmt.Sprintf(
			"%s%s\n%s%s",
			indent,
			"chair tldr alp",
			indent,
			"chair tldr pt-query-digest",
		),
		RunE: run,
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		switch args[0] {
		case "alp":
			fmt.Println("sudo cat /var/log/nginx/access.log | alp -r")
		case "pt-query-digest":
			fmt.Println("sudo pt-query-digest /var/log/mysql/mysqld-slow.log")
		default:
			return cmd.Usage()
		}
		return nil
	}
	return cmd.Usage()
}
