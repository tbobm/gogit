package main

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/google/go-github/github"
    "context"
)

func createRepo(c *cli.Context) error {
    client := getClient()
    fmt.Println(c.Bool("private"))

    name := c.String("name")
    private := c.Bool("private")
    if name == "" {
        return fmt.Errorf("Missing repository name")
    }

    repo := &github.Repository{
        Name: github.String(name),
        Private: github.Bool(private),
    }

    repository, response, err := client.Repositories.Create(context.Background(), "", repo)
    fmt.Println(repository, response, err)
    return nil
}


func listRepos(c *cli.Context) {
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
    for idx, repo := range repos {
        fmt.Println(repo.GetName(), ": ", repo.GetURL())
        if idx == maxDisplay {
            break
        }
    }
}

