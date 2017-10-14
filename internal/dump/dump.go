package dump

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const indent = "    "

type dumper struct {
	port     int
	hostname string
	username string
	password string
	database string
}

func CommandNew() *cobra.Command {
	d := new(dumper)
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "Dump mysql schema or data",
		Long:  "Dump mysql schema or data",
		Example: fmt.Sprintf(
			"%s%s\n%s%s\n%s%s",
			indent,
			"chair dump [flags] all",
			indent,
			"chair dump [flags] schema",
			indent,
			"chair dump [flags] data",
		),
		RunE: d.run,
		ValidArgs: []string{
			"all",
			"schema",
			"data",
		},
	}

	cmd.Flags().IntVarP(&d.port, "port", "P", 3306, "specify mysql port number")
	cmd.Flags().StringVarP(&d.hostname, "host", "H", "127.0.0.1", "specify mysql hostname")
	cmd.Flags().StringVarP(&d.username, "user", "u", "", "specify mysql username")
	cmd.Flags().StringVarP(&d.password, "pass", "p", "", "specify mysql password")
	cmd.Flags().StringVarP(&d.database, "database", "d", "", "specify mysql database")

	return cmd
}

func (d *dumper) run(c *cobra.Command, args []string) error {
	if len(args) > 0 {
		cmdArgs := make([]string, 0, 2)
		cmdArgs = append(cmdArgs, "-P"+fmt.Sprintf("%d", d.port))
		cmdArgs = append(cmdArgs, "-h"+d.hostname)
		if d.username != "" {
			cmdArgs = append(cmdArgs, "-u"+d.username)
		}
		if d.password != "" {
			cmdArgs = append(cmdArgs, "-p"+d.password)
		}
		switch args[0] {
		case "schema":
			cmdArgs = append(cmdArgs, "--no-data")
		case "data":
			cmdArgs = append(cmdArgs, "--no-create-info")
		case "all":
		default:
			return errors.New("Please specify which one: schema, data, all")
		}
		if d.database != "" {
			cmdArgs = append(cmdArgs, d.database)
		}

		cmd := exec.Command("mysqldump", cmdArgs...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "Failed to create pipe to stdout")
		}
		if err := cmd.Start(); err != nil {
			return errors.Wrap(err, "Failed to run mysqldump")
		}
		now := time.Now().Format("20060102_150405")
		f, err := os.Create(
			fmt.Sprintf("%s_%s.sql", args[0], now),
		)
		if err != nil {
			return errors.Wrap(err, "Failed to create sql file")
		}
		f.Chmod(0644)
		if _, err := io.Copy(f, stdout); err != nil {
			return errors.Wrap(err, "Failed to dump sql file")
		}
		return nil
	}
	return c.Usage()
}
