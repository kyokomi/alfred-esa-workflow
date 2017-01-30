package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/pascalw/go-alfred"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// SearchService posts search service
type SearchService struct {
	*Workflow
}

// Command cli command
func (s *SearchService) Command(c *cli.Context) {
	if !s.Config.IsValid() {
		s.Alfred.PrintError(errors.New("don't setup config"))
		os.Exit(1)
	}

	if err := s.run(c.Args()); err != nil {
		s.Alfred.PrintError(err)
		os.Exit(1)
	}
}

func (s *SearchService) run(args []string) error {
	query := url.Values{}
	for _, arg := range args {
		query.Add("", arg)
	}

	resp, err := s.Client.Post.GetPosts(s.Config.TeamName, query)
	if err != nil {
		return err
	}

	items := make([]*alfred.AlfredResponseItem, len(resp.Posts))
	for i, post := range resp.Posts {
		items[i] = &alfred.AlfredResponseItem{
			Valid:    true,
			Uid:      strconv.Itoa(post.Number),
			Title:    post.Name,
			Arg:      post.URL,
			Subtitle: fmt.Sprintf("%s %s", post.Category, post.CreatedAt),
		}
	}

	s.Alfred.PrintItems(items)
	return nil
}
