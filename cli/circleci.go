package cli

import (
	"fmt"

	"github.com/KeisukeYamashita/biko/browser"
	cc "github.com/KeisukeYamashita/biko/providers/circleci"
	"github.com/urfave/cli"
)

func newCircleCICmd() cli.Command {
	return cli.Command{
		Name:    "circleci",
		Aliases: []string{"cc"},
		Usage:   "Open CircleCI resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_CIRCLECI",
				Usage:  "Specify CircleCI Organization",
			},
		},
		Action: func(c *cli.Context) error {
			var org string
			if org = c.String("org"); org == "" {
				return fmt.Errorf("Org for circleci not configured pass --org or set BIKO_CIRCLECI")
			}
			circleci := &cc.Provider{
				Org: org,
			}
			return browser.Open(c, circleci)
		},
		Subcommands: []cli.Command{
			newCircleCIJobsCmd(),
		},
	}
}

func newCircleCIJobsCmd() cli.Command {
	return cli.Command{
		Name:    "jobs",
		Aliases: []string{"j"},
		Usage:   "Open incident page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_PAGERDUTY",
				Usage:  "Specify Pagerduty Organization",
			},
		},
		Action: func(c *cli.Context) error {
			cc, err := cc.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, cc)
		},
	}
}
