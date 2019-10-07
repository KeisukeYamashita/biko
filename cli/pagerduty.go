package cli

import (
	"fmt"

	"github.com/KeisukeYamashita/biko/browser"
	pd "github.com/KeisukeYamashita/biko/providers/pagerduty"
	"github.com/urfave/cli"
)

func newPagerDutyCmd() cli.Command {
	return cli.Command{
		Name:    "pagerduty",
		Aliases: []string{"pd"},
		Usage:   "Open PagerDuty resource",
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
			var org string
			if org = c.String("org"); org == "" {
				return fmt.Errorf("Org for pagerduty not configured pass --org or set BIKO_PAGERDUTY")
			}
			pd := &pd.Provider{
				Org: org,
			}
			return browser.Open(c, pd)
		},
		Subcommands: []cli.Command{
			newPagerDudyIncidentCmd(),
			newPagerDudyAlertCmd(),
			newPagerDudySchedulesCmd(),
		},
	}
}

func newPagerDudyIncidentCmd() cli.Command {
	return cli.Command{
		Name:    "incident",
		Aliases: []string{"i"},
		Usage:   "Open incident page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_PAGERDUTY",
				Usage:  "Specify Pagerduty Organization",
			},
		},
		Action: func(c *cli.Context) error {
			var org string
			if org = c.String("org"); org == "" {
				return fmt.Errorf("Org for pagerduty not configured pass --org or set BIKO_PAGERDUTY")
			}
			pd := &pd.Provider{
				Org: org,
			}
			return browser.Open(c, pd)
		},
	}

}

func newPagerDudyAlertCmd() cli.Command {
	return cli.Command{
		Name:    "alert",
		Aliases: []string{"a"},
		Usage:   "Open alert page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			var org string
			if org = c.String("org"); org == "" {
				return fmt.Errorf("Org for pagerduty not configured pass --org or set BIKO_PAGERDUTY")
			}
			pd := &pd.Provider{
				Org: org,
			}
			return browser.Open(c, pd)
		},
	}
}

func newPagerDudySchedulesCmd() cli.Command {
	return cli.Command{
		Name:    "schedules",
		Aliases: []string{"s"},
		Usage:   "Open schedules page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_PAGERDUTY",
				Usage:  "Specify Pagerduty Organization",
			},
		},
		Action: func(c *cli.Context) error {
			var org string
			if org = c.String("org"); org == "" {
				return fmt.Errorf("Org for pagerduty not configured pass --org or set BIKO_PAGERDUTY")
			}
			pd := &pd.Provider{
				Org: org,
			}
			return browser.Open(c, pd)
		},
	}
}
