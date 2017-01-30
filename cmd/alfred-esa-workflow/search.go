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

func cmdSearch(c *cli.Context) {
	w, err := newEsaWorkflowFromConfig()
	if err != nil {
		alfredPrintError(errors.New("don't setup config"))
		os.Exit(1)
	}

	items, err := w.searchPosts(c.Args())
	if err != nil {
		alfredPrintError(err)
		os.Exit(1)
	}

	alfredPrintItems(items)
}

func (w *esaWorkflow) searchPosts(args []string) ([]*alfred.AlfredResponseItem, error) {
	query := url.Values{}
	for _, arg := range args {
		query.Add("", arg)
	}

	resp, err := w.client.Post.GetPosts(w.teamName, query)
	if err != nil {
		return nil, err
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
	return items, nil
}
