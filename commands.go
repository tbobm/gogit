package main

import (
    //"fmt"
    "github.com/urfave/cli"
    //"context"
)

func commandsCreateRepo() cli.Command {
    command := &cli.Command{
        Name:  "create",
        Aliases: []string{"c"},
        Usage: "create a new repository",
        Action: createRepo,
        Flags: []cli.Flag {
            cli.StringFlag{
                Name: "name,n",
                Usage: "repository name",
            },
            cli.BoolFlag{
                Name: "private,p",
                Usage: "repository's visibility",
            },
        },
    }
    return *command
}

func commandsListRepo() cli.Command {
    command := &cli.Command{
        Name: "list",
        Aliases: []string{"l"},
        Usage: "list top ten last updated repositories",
        Action: listRepos,
    }
    command.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "user,u",
            Usage: "the author of the repositories",
        },
        cli.Int64Flag{
            Name: "count,c",
            Usage: "number of repositories to display",
        },
    }
    return *command
}
