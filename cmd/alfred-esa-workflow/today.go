package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

// TodayService today posts get service
type TodayService struct {
	*Workflow
}

// Command cli command
func (s *TodayService) Command(c *cli.Context) {
	if !s.Config.IsValid() {
		s.Alfred.PrintError(errors.New("don't setup config"))
		os.Exit(1)
	}

	if err := s.run(c.Args()); err != nil {
		s.Alfred.PrintError(err)
		os.Exit(1)
	}
}

func (s *TodayService) run(args []string) error {
	query := url.Values{}
	if len(args) < 1 {
		return createArgsError(args)
	}
	username := args[0]
	today := nowTimeTodayDate().Format("2006-01-02")
	if len(args) > 1 {
		if inputDate := argTimeTodayDate(args[1]); !inputDate.IsZero() {
			today = inputDate.Format("2006-01-02")
		}
	}

	query.Add("updated", fmt.Sprintf(">=%s", today))
	query.Add("user", username)

	resp, err := s.Client.Post.GetPosts(s.Config.TeamName, query)
	if err != nil {
		return err
	}

	fullNameURLs := make([]string, len(resp.Posts))
	for i, post := range resp.Posts {
		if post.Wip {
			fullNameURLs[i] = fmt.Sprintf("- [ ] [#%d: [WIP] %s](%s)", post.Number, post.FullName, post.URL)
		} else {
			fullNameURLs[i] = fmt.Sprintf("- [x] [#%d: %s](%s)", post.Number, post.FullName, post.URL)
		}
	}

	if len(fullNameURLs) > 0 {
		fmt.Println(strings.Join(fullNameURLs, "\n"))
		return nil
	}

	fmt.Println("ERROR: there are no posts of today")
	return nil
}

func nowTimeTodayDate() time.Time {
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
}

func createArgsError(args []string) error {
	return fmt.Errorf("args error: <username> <yyyy-MM-dd | number of hours ago> != %s", args)
}

func argTimeTodayDate(arg string) time.Time {
	inputHoursAgo, err := strconv.Atoi(arg)
	if err == nil {
		return time.Now().Add(time.Hour * time.Duration(inputHoursAgo*-1))
	}
	inputDate, err := time.Parse("2006-01-02", arg)
	if err == nil {
		return inputDate
	}
	return time.Time{}
}
