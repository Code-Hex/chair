package main

import (
	"github.com/Code-Hex/chair"
)

func main() {
	c := chair.NewCommand()
	err := c.GenBashCompletionFile("chair.bash")
	if err != nil {
		panic(err)
	}
}
