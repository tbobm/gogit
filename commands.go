package main

import (
	//"fmt"
	"github.com/urfave/cli"
	//"context"
)

func commandsCreateRepo() cli.Command {
	command := &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create a new repository",
		Action:  createRepo,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name,n",
				Usage: "repository name",
			},
			cli.BoolFlag{
				Name:  "private,p",
				Usage: "repository's visibility",
			},
		},
	}
	return *command
}

func commandsAddCollab() cli.Command {
	command := &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a collaborator",
		Action:  addCollab,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name,n",
				Usage: "repository name",
			},
			cli.StringFlag{
				Name:  "owner,o",
				Usage: "repository owner name",
			},
			cli.StringFlag{
				Name:  "collaborator,c",
				Usage: "collaborator's username",
			},
		},
	}
	return *command
}

func commandsListRepo() cli.Command {
	command := &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list top ten last updated repositories",
		Action:  listRepos,
	}
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username,user,u",
			Usage: "the author of the repositories",
		},
		cli.Int64Flag{
			Name:  "count,c",
			Usage: "number of repositories to display",
		},
	}
	return *command
}
