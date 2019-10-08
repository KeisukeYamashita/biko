package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	"github.com/KeisukeYamashita/biko/providers/google"
	"github.com/urfave/cli"
)

func newGoogleCmd() cli.Command {
	return cli.Command{
		Name:    "google",
		Aliases: []string{"g"},
		Usage:   "Open Google source",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := google.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
		Subcommands: []cli.Command{
			newGoogleSearchCmd(),
		},
	}
}

func newGoogleSearchCmd() cli.Command {
	return cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "Search a page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := google.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}
