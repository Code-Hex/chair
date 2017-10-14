package tldr

import (
	"bytes"
	"os"

	"github.com/fatih/color"
)

type alp struct {
	name string
}

func alpNew() TLDR {
	return &alp{
		name: "alp",
	}
}

func (a *alp) Print() {
	buf := bytes.NewBufferString("\nalp is Access Log Profiler for Labeled Tab-separated Values(LTSV).\n\n")
	buf.WriteString(color.GreenString("- Read access log") + "\n")
	buf.WriteString(indent + color.RedString("alp -f") + " " + color.BlueString("/var/log/nginx/access.log") + "\n\n")
	buf.WriteString(color.GreenString("- Include query and reverse") + "\n")
	buf.WriteString(indent + color.RedString("alp -f") + " " + color.BlueString("/var/log/nginx/access.log") + " ")
	buf.WriteString(color.BlueString("-q -r") + "\n\n")
	buf.WriteString(color.GreenString("- Include status code") + "\n")
	buf.WriteString(indent + color.RedString("cat") + " " + color.BlueString("/var/log/nginx/access.log") + color.RedString(" | "))
	buf.WriteString(color.RedString("alp --include-statuses ") + color.BlueString("200") + "\n\n")
	buf.WriteString(color.GreenString("- Exclude status code") + "\n")
	buf.WriteString(indent + color.RedString("cat") + " " + color.BlueString("/var/log/nginx/access.log") + color.RedString(" | "))
	buf.WriteString(color.RedString("alp --exclude-statuses ") + color.BlueString("20[0-9]") + "\n\n")
	buf.WriteString(color.GreenString("- Aggregate URI") + "\n")
	buf.WriteString(indent + color.RedString("alp -f") + " " + color.BlueString("/var/log/nginx/access.log") + " ")
	buf.WriteString(color.RedString("--aggregates ") + color.BlueString(`"/diary/entry/\d+"`) + "\n\n")
	os.Stdout.Write(buf.Bytes())
}

func (a *alp) Name() string           { return a.name }
func (a *alp) Match(name string) bool { return a.name == name }
