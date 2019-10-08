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
			github := &gh.Provider{}
			return browser.Open(c, github)
		},
		Subcommands: []cli.Command{
			newGithubDashboardCmd(),
			newGithubTrendingCmd(),
		},
	}
}

func newGithubDashboardCmd() cli.Command {
	return cli.Command{
		Name:    "dashboard",
		Aliases: []string{"db"},
		Usage:   "Open a dashboard page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "org",
				Usage: "Organization to open",
			},
		},
		Action: func(c *cli.Context) error {

			github := &gh.Provider{}
			return browser.Open(c, github)
		},
	}

}

func newGithubTrendingCmd() cli.Command {
	return cli.Command{
		Name:    "trending",
		Aliases: []string{"t"},
		Usage:   "Open a trending page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "language, l",
				Usage: "filter trending with language",
			},
			cli.StringFlag{
				Name:  "since, s",
				Usage: "filter trending with date range(daily, weekly, monthly)",
			},
		},
		Action: func(c *cli.Context) error {

			github := &gh.Provider{}
			return browser.Open(c, github)
		},
	}

}
