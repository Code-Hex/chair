package main

import (
	"fmt"
	"os"

	"github.com/Code-Hex/chair/dump"
	"github.com/pkg/errors"
)

func (c *Chair) runShow() error {
	switch c.Show {
	// MySQL
	case "sql-default":
		fmt.Println(dump.SQLDefaultConfig())
	case "sql-slow-log":
		fmt.Println(dump.SQLSlowLogConfig())
	case "sql-maybe-nice":
		dumpStr, err := dump.SQLMaybeNiceConfig()
		if err != nil {
			return errors.Wrap(err, "Failed to get memory size")
		}
		fmt.Println(dumpStr)
	case "sql-cache":
		fmt.Println(dump.SQLCacheConfig())
	case "sql-fix57-groupby":
		fmt.Println(dump.SQLFix57GroupByProblem())
	// Nginx
	case "nginx-access-log":
		fmt.Println(dump.NginxAccessLogConfig())
	case "nginx-event":
		fmt.Println(dump.NginxEventConfig())
	case "nginx-outside-maybe-nice":
		fmt.Println(dump.NginxOutsideMaybeNiceConfig())
	case "nginx-static":
		fmt.Println(dump.NginxStaticLocationConfig())
	default:
		os.Stderr.Write(c.usage())
		return nil
	}
	return nil
}
