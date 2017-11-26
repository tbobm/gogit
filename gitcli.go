package main

import (
	"fmt"
    "net/http"
	"github.com/urfave/cli"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    "context"
	//"log"
	"os"
)

const maxDisplay int = 10

func createRepo(name string, private bool) error {
    //client := github.NewClient(nil)

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


func getClient() *github.Client {
    var token string
    var tc *http.Client

    token = os.Getenv("GIT_TOKEN")
    if token != "" {
        ts := oauth2.StaticTokenSource(
            &oauth2.Token{AccessToken: token},
        )
        tc = oauth2.NewClient(context.Background(), ts)
    }

    return github.NewClient(tc)
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Repository management",
			Subcommands: []cli.Command{
				{
					Name:  "Create",
					Usage: "create a new repository",
					Action: func(c *cli.Context) error {
						fmt.Println("New repository created", c.Args().First())
						return nil
					},
                }, {
                    Name: "List",
                    Usage: "list top ten last updated repositories",
                    Action: listRepos,
                },
			},
		},
	}
	app.Run(os.Args)
    //listRepos(client)
}
