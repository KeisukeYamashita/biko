// Copyright 2019 The Biko Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/KeisukeYamashita/biko/browser"
	pd "github.com/KeisukeYamashita/biko/providers/pagerduty"
	"github.com/urfave/cli"
)

func newPagerDutyCmd() cli.Command {
	return cli.Command{
		Name:     "pagerduty",
		Aliases:  []string{"pd"},
		Usage:    "Open PagerDuty resource",
		Category: categoryIncidentManagement,
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
