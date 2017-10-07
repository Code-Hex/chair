package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/Code-Hex/chair/dump"
	"github.com/Code-Hex/chair/gen"
	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

const (
	version = "0.0.1"
	name    = "chair"
	msg     = name + " is initial setup tool for isucon"
)

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
	if os.Geteuid() > 0 {
		return errors.New("You must run this program as a superuser")
	}
	_, err := c.prepare()
	if err != nil {
		return err
	}

	if c.Dump != "" {
		var dumpStr string
		switch c.Dump {
		// MySQL
		case "sql-default":
			dumpStr = dump.SQLDefaultConfig()
		case "sql-slow-log":
			dumpStr = dump.SQLSlowLogConfig()
		case "sql-maybe-nice":
			var err error
			dumpStr, err = dump.SQLMaybeNiceConfig()
			if err != nil {
				return errors.Wrap(err, "Failed to get memory size")
			}
		case "sql-cache":
			dumpStr = dump.SQLCacheConfig()
		// Nginx
		case "nginx-access-log":
			dumpStr = dump.NginxAccessLogConfig()
		case "nginx-event":
			dumpStr = dump.NginxEventConfig()
		case "nginx-outside-maybe-nice":
			dumpStr = dump.NginxOutsideMaybeNiceConfig()
		case "nginx-static":
			dumpStr = dump.NginxStaticLocationConfig()
		default:
			os.Stderr.Write(c.usage())
			return nil
		}
		fmt.Println(dumpStr)
		return nil
	}

	if c.Init {
		restartScript := gen.GenerateRestartScript()
		f, err := os.Create("restart.sh")
		if err != nil {
			return errors.Wrap(err, "Failed to create restart.sh")
		}
		f.WriteString(restartScript)
		f.Chmod(0755)
		f.Close()

		files := []file{
			alpNew(),
			ptQueryDigestNew(),
		}

		var wg sync.WaitGroup
		for _, v := range files {
			fmt.Println("Start:", v.URL())
			wg.Add(1)
			go func(fi file) {
				defer wg.Done()
				if err := download(fi); err != nil {
					panic(err)
				}
				fmt.Println("Done:", fi.URL())
			}(v)
		}
		wg.Wait()
		return nil
	}

	os.Stderr.Write(c.usage())
	return nil
}

func download(file file) error {
	req, err := http.NewRequest("GET", file.URL(), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to make request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Failed to download: %s", file.URL())
	}
	defer resp.Body.Close()

	f, err := os.Create(file.Name())
	if err != nil {
		return errors.Wrapf(err, "Failed to create %s", file.Name())
	}
	io.Copy(f, resp.Body)
	f.Chmod(0755)
	f.Close()

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
