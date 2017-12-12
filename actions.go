package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

func createRepo(c *cli.Context) error {
	verbose := true
	client := getClient()
	fmt.Println(c.Bool("private"))

	fmt.Println("Verbose: ", c.Bool("verbose"))
	name := c.String("name")
	private := c.Bool("private")
	if name == "" {
		return fmt.Errorf("Missing repository name")
	}

	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	}

	repository, response, err := client.Repositories.Create(context.Background(), "", repo)
	fmt.Println("Created repository: ", repository.GetName())
	fmt.Println("URL: ", repository.GetURL())

	if verbose {
		fmt.Println("Remaining calls: ", response.Remaining)
		fmt.Println("Reset timestamp: ", response.Rate.Reset)
	}

	if err != nil {
		return err
	}
	return nil
}

func listRepos(c *cli.Context) error {
	var toDisplay int64
	client := getClient()
	opt := &github.RepositoryListOptions{
		Sort: "updated",
	}
	uname := c.String("user")
	if uname != "" {
		fmt.Println("Listing repositories for user: ", uname)
	}

	repos, resp, err := client.Repositories.List(context.Background(), uname, opt)
	if err != nil {
		return err
	}
	argDisplay := c.Int64("count")

	if argDisplay == 0 {
		toDisplay = maxDisplay
	} else {
		toDisplay = argDisplay
	}

	var idxAs64 int64
	for idx, repo := range repos {
		idxAs64 = int64(idx + 1)
		fmt.Println(idx+1, ":", repo.GetName(), "-", repo.GetHTMLURL())
		if idxAs64 == toDisplay {
			break
		}
	}
	fmt.Println(resp.Rate.Remaining)
	return nil
}
