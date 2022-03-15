package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func main() {
	// Use 4 command line arguments: owner, repo, token, issue ,title, label and body
	if len(os.Args) != 7 {
		fmt.Println(len(os.Args))
		panic("Usage: main.go owner repo token title label and body")
	}
	// parse command line arguments
	owner := os.Args[1]
	repo := os.Args[2]
	token := os.Args[3]
	title := os.Args[4]
	label := os.Args[5]
	body := os.Args[6]
	labels := strings.Split(label, ",")
	// validate command line arguments
	if owner == "" || repo == "" || token == "" || title == "" || label == "" {
		panic("Usage: main.go owner repo token title and label")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	var issue *github.Issue
	var err error
	issue, _, err = client.Issues.Create(ctx, owner, repo, &github.IssueRequest{
		Title:  github.String(title),
		Body:   github.String(body),
		Labels: &labels,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("Issue %d created in the repository %s", *issue.Number, *issue.RepositoryURL))
}
