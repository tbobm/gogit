package main

import (
    "fmt"
    "github.com/urfave/cli"
    //"context"
)

func commandsCreateRepo() cli.Command {
    command := &cli.Command{
        Name:  "create",
        Aliases: []string{"c"},
        Usage: "create a new repository",
        Action: func(c *cli.Context) error {
            fmt.Println("New repository created", c.Args().First())
            return nil
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
    return *command
}
