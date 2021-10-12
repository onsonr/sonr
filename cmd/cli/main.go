package main

import (
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Name: "sonr",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language for the greeting",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
