package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func cmdSetup(c *cli.Context) {
	args := c.Args().First()
	if err := setup(strings.Fields(args)); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("saved config")
}

func setup(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("args error: <accessToken> <teamName> != %d", len(args))
	}
	accessToken := args[0]
	teamName := args[1]

	if err := saveConfig(config{AccessToken: accessToken, TeamName: teamName}); err != nil {
		return fmt.Errorf("config save error %s", err.Error())
	}

	return nil
}
