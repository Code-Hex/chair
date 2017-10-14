<p align="center">
  <img alt="logo" src="https://user-images.githubusercontent.com/6500104/31577054-c9f68064-b142-11e7-9e5a-e41046f330f7.png" height="140" />
  <h3 align="center">Chair</h3>
  <p align="center">This chair help you when you play isucon.</p>
  <p align="center">
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
  </p>
</p>

---

# Usage

```
Usage:
  chair [flags]
  chair [command]

Available Commands:
  dump        Dump mysql schema or data
  help        Help about any command
  info        show machine information
  setup       Run initial setup for isucon environment
  show        Show some configuration content
  tldr        tldr some command

Flags:
  -h, --help      help for chair
      --version   show version (default true)
```

## SubCommands
### dump
Dump mysql schema or data.
```
Usage:
  chair dump [flags]

Examples:
    chair dump [flags] all
    chair dump [flags] schema
    chair dump [flags] data

Flags:
  -d, --database string   specify mysql database
  -h, --help              help for dump
  -H, --host string       specify mysql hostname (default "127.0.0.1")
  -p, --pass string       specify mysql password
  -P, --port int          specify mysql port number (default 3306)
  -u, --user string       specify mysql username
```

### info
Check information you are using machine.
```
Usage:
  chair info [flags]

Flags:
  -h, --help            help for info
  -t, --time duration   specify time to measure cpu utilization (default 1s)
```

### setup
`setup` command will download some binary, generate script to help you.  
If you want to run, you should run as a superuser.
```
Usage:
  chair setup [flags]

Examples:
    sudo chair setup

Flags:
  -h, --help   help for setup
```

### show
Show a part of configuration for mysql and nginx.
```
Usage:
  chair show [flags]

Examples:
    chair show sql-default
    chair show nginx-access-log
    chair show sql-(default|slow-log|maybe-nice|cache|fix57-groupby)
    chair show nginx-(access-log|event|outside-maybe-nice|static)
```

### tldr
Show tldr for some command to help you.
```
Usage:
  chair tldr [flags]

Examples:
    chair tldr alp
    chair tldr pt-query-digest

Flags:
  -h, --help   help for tldr
```
