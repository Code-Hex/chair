package tldr

import (
	"bytes"
	"os"

	"github.com/fatih/color"
)

type ptQueryDigest struct {
	name string
}

func ptQueryDigestNew() TLDR {
	return &ptQueryDigest{
		name: "pt-query-digest",
	}
}

func (p *ptQueryDigest) Print() {
	buf := bytes.NewBufferString("\nAnalyze MySQL queries from logs, processlist, and tcpdump.\n\n")
	buf.WriteString(color.GreenString("- Report the slowest queries from") + "\n")
	buf.WriteString(indent + color.RedString("pt-query-digest /var/log/mysql/mysqld-slow.log") + "\n\n")
	os.Stdout.Write(buf.Bytes())
}

func (p *ptQueryDigest) Name() string           { return p.name }
func (p *ptQueryDigest) Match(name string) bool { return p.name == name }
