package main

import (
	"bytes"
	"fmt"
	"os"

	"reflect"

	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

const indent = "        "

// Options struct for parse command line arguments
type Options struct {
	Help    string `short:"h" long:"help" description:"args: alp, pt-query-digest"`
	Version bool   `short:"v" long:"version" description:"print the version"`
	Init    bool   `short:"i" long:"init" description:"initialize"`

	Dump       string `short:"d" long:"dump" description:"args: sql-(default|slow-log|maybe-nice|cache), nginx-(access-log|event|outside-maybe-nice|static)"`
	StackTrace bool   `long:"trace" description:"display detail error messages"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	p := flags.NewParser(opts, flags.None)
	args, err := p.ParseArgs(argv)
	if err != nil {
		os.Stderr.Write(opts.usage())
		return nil, errors.Wrap(err, "invalid command line options")
	}
	return args, nil
}

func (opts Options) usage() []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `%s
Usage: %s [options]
Options:
`, msg, name)

	t := reflect.TypeOf(opts)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		desc := tag.Get("description")
		var o string
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s", tag.Get("short"), tag.Get("long"))
		} else {
			o = fmt.Sprintf("--%s", tag.Get("long"))
		}
		fmt.Fprintf(&buf, "  %-21s %s\n", o, desc)

		if deflt := tag.Get("default"); deflt != "" {
			fmt.Fprintf(&buf, "  %-21s default: --%s='%s'\n", indent, tag.Get("long"), deflt)
		}
	}

	return buf.Bytes()
}
