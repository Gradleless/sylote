package main

import (
	"os"

	"github.com/gradleless/sylote/lib"
	"github.com/gradleless/sylote/lib/cli"
)

func main() {
	if lib.CheckEnable() && len(os.Args) == 1 {
		cli.Automation()
		return
	}
	cli.Cli(os.Args)
}
