package main

import (
	"os"

	"github.com/mavenless/sylote/lib"
	"github.com/mavenless/sylote/lib/cli"
)

func main() {
	if lib.CheckEnable() && len(os.Args) == 1 {
		cli.Automation()
		return
	}
	cli.Cli(os.Args)
}
