package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	bundleID = "com.kyokomi.alfred-esa-workflow"
	version  = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "alfred-esa-workflow"
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Action: cmdSetup,
		},
		{
			Name:   "search",
			Action: cmdSearch,
		},
	}
	app.Run(os.Args)
}
