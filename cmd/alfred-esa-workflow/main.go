package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	w := NewWorkflow()

	app := cli.NewApp()
	app.Name = w.AppName
	app.Version = w.Version
	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Action: w.Setup.Command,
		},
		{
			Name:   "search",
			Action: w.Search.Command,
		},
		{
			Name:   "today",
			Action: w.Today.Command,
		},
	}
	app.Run(os.Args)
}
