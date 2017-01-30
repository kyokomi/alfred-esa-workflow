package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// SetupService setup service
type SetupService struct {
	*Workflow
}

// Command cli command
func (s *SetupService) Command(c *cli.Context) {
	args := c.Args().First()
	if err := s.run(strings.Fields(args)); err != nil {
		s.Alfred.Error(err)
		os.Exit(1)
	}
	s.Alfred.Message("saved config")
}

func (s *SetupService) run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("args error: <accessToken> <teamName> != %d", len(args))
	}
	accessToken := args[0]
	teamName := args[1]

	s.Config = Config{AccessToken: accessToken, TeamName: teamName}
	if err := s.SaveConfig(); err != nil {
		return fmt.Errorf("config save error %s", err.Error())
	}

	return nil
}
