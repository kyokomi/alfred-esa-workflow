package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	w := NewWorkflow()

	app := cli.NewApp()
	app.Name = os.Getenv("alfred_workflow_bundleid")
	app.Version = os.Getenv("alfred_version")
	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Action: w.Setup.Command,
		},
		{
			Name:   "search",
			Action: w.Search.Command,
		},
	}
	app.Run(os.Args)
}
