package main

import (
	"github.com/urfave/cli"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    "context"
    "net/http"
	//"log"
	"os"
)

const maxDisplay int = 10

func getClient() *github.Client {
    var token string
    var tc *http.Client

    token = os.Getenv("GITHUB_TOKEN")
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
				commandsCreateRepo(), commandsListRepo(),
			},
		},
	}
	app.Run(os.Args)
    //listRepos(client)
}
