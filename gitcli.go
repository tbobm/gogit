package main

import (
	"fmt"
	"github.com/urfave/cli"
    "github.com/google/go-github"
	"log"
	"os"
)

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
				},
			},
		},
	}
	log.Println("Gitcli")
	app.Run(os.Args)
}
