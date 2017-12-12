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

func listRepos(c *cli.Context) {
	var toDisplay int64
	client := getClient()
	opt := &github.RepositoryListOptions{
		Sort: "updated",
	}
	repos, resp, err := client.Repositories.List(context.Background(), "tbobm", opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Rate.Remaining)
	fmt.Println("10 last updated repositories: ")
	argDisplay := c.Int64("count")

	fmt.Println("Count:", argDisplay)
	if argDisplay == 0 {
		toDisplay = maxDisplay
	} else {
		toDisplay = argDisplay
	}

	var idxAs64 int64
	for idx, repo := range repos {
		idxAs64 = int64(idx)
		fmt.Println(repo.GetName(), ": ", repo.GetURL())
		if idxAs64 == toDisplay {
			break
		}
	}
}
