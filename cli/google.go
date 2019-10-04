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
		Usage:   "Open GCP resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &google.Provider{}
			return browser.Open(c, gcp)
		},
		Subcommands: []cli.Command{
			newSearchCmd(),
		},
	}
}

func newSearchCmd() cli.Command {
	return cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "Search a page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query",
				Usage: "Query a page",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &google.Provider{}
			return browser.Open(c, gcp)
		},
	}

}
