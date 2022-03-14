package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func main() {
	// Use 4 command line arguments: owner, repo, token, issue title and label
	if len(os.Args) != 6 {
		fmt.Println(len(os.Args))
		panic("Usage: main.go owner repo token title and label")
	}
	// parse command line arguments
	owner := os.Args[1]
	repo := os.Args[2]
	token := os.Args[3]
	title := os.Args[4]
	label := os.Args[5]
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
		Title: github.String(title),
		Labels: &[]string{
			label,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("Issue %d created in the repository %s", *issue.Number, *issue.RepositoryURL))
}
