package main

import (
	"github.com/urfave/cli"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    "context"
    "net/http"
	"os"
)

const maxDisplay int64 = 10

const (
    author string = "Theo Massard"
    email string = "<massar_t@etna-alternance.net>"
    version string = "1.1.0"
)

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

    app.Name = "Gogit"
    app.Author = author
    app.Version = version
    app.Usage = "CLI client for github.com"
    app.Email = email

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
}
