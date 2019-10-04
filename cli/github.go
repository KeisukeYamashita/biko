package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	gh "github.com/KeisukeYamashita/biko/providers/github"
	"github.com/urfave/cli"
)

func newGithubCmd() cli.Command {
	return cli.Command{
		Name:    "github",
		Aliases: []string{"gh"},
		Usage:   "Open Github resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gh.Provider{}
			return browser.Open(c, gcp)
		},
		Subcommands: []cli.Command{
			newGithubOrgCmd(),
		},
	}
}

func newGithubOrgCmd() cli.Command {
	return cli.Command{
		Name:    "organization",
		Aliases: []string{"org"},
		Usage:   "Open a organization page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gh.Provider{}
			return browser.Open(c, gcp)
		},
	}

}
