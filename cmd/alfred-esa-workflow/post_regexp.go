package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// PostRegexpService posts hoge service
type PostRegexpService struct {
	*Workflow
}

type params struct {
	printPrefix string
	regexp      *regexp.Regexp
	query       url.Values
}

// Command cli command
func (s *PostRegexpService) Command(c *cli.Context) {
	if !s.Config.IsValid() {
		s.Alfred.PrintError(errors.New("don't setup config"))
		os.Exit(1)
	}

	if err := s.run(c.Args()); err != nil {
		s.Alfred.PrintError(err)
		os.Exit(1)
	}
}

func (PostRegexpService) parseParams(args []string) (params, error) {
	if len(args) < 3 {
		return params{}, createArgsError(args)
	}
	q := url.Values{}
	q.Add("user", args[0])
	today := time.Date(2017, time.July, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	q.Add("created", fmt.Sprintf(">=%s", today))
	q.Add("keyword", args[1])

	result := params{
		query:  q,
		regexp: regexp.MustCompile(args[2]),
	}

	if len(args) > 3 {
		result.printPrefix = args[3]
	}

	return result, nil
}

func (s *PostRegexpService) run(args []string) error {
	ps, err := s.parseParams(args)
	if err != nil {
		return err
	}
	// TODO: ページネーション対応が保留
	// https://docs.esa.io/posts/102#2-7-0
	resp, err := s.Client.Post.GetPosts(s.Config.TeamName, ps.query)
	if err != nil {
		return err
	}

	var printText []string
	for _, post := range resp.Posts {
		keywords := ps.regexp.FindSubmatch([]byte(post.BodyMd))
		if len(keywords) <= 1 {
			continue
		}

		printText = append(printText, ps.printPrefix+post.FullName)
		printText = append(printText, string(keywords[1]))
	}
	fmt.Println(strings.Join(printText, ""))

	return nil
}

func (PostRegexpService) createArgsError(args []string) error {
	return fmt.Errorf("args error: <username> <keyword> <regexp> <printPrefix (optional)> != %s", args)
}
