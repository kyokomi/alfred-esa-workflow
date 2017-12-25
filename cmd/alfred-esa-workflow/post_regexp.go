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

	printPrefix string
	regexp      *regexp.Regexp
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

func (s *PostRegexpService) run(args []string) error {
	query := url.Values{}
	if len(args) < 1 {
		return createArgsError(args)
	}
	username := args[0]
	query.Add("user", username)
	today := time.Date(2017, time.July, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	query.Add("created", fmt.Sprintf(">=%s", today))
	query.Add("keyword", "日報")
	if len(args) > 1 {
		query.Set("keyword", args[1])
	}

	s.printPrefix = "## "
	if len(args) > 2 {
		s.printPrefix = args[2]
	}
	s.regexp = regexp.MustCompile(`# 悪かったこと([\s\S]*)# 所感`)
	if len(args) > 3 {
		s.regexp = regexp.MustCompile(args[3])
	}

	// TODO: ページネーション対応が保留
	// https://docs.esa.io/posts/102#2-7-0
	resp, err := s.Client.Post.GetPosts(s.Config.TeamName, query)
	if err != nil {
		return err
	}

	var printText []string
	for _, post := range resp.Posts {
		keywords := s.regexp.FindSubmatch([]byte(post.BodyMd))
		if len(keywords) <= 1 {
			continue
		}

		printText = append(printText, s.printPrefix+post.FullName)
		printText = append(printText, string(keywords[1]))
	}
	fmt.Println(strings.Join(printText, ""))

	return nil
}
