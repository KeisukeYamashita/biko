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
			cc := &cc.Provider{
				Org: org,
			}
			return browser.Open(c, cc)
		},
		Subcommands: []cli.Command{
			newCircleCIJobsCmd(),
			newCircleCIWorkflowsCmd(),
		},
	}
}

func newCircleCIJobsCmd() cli.Command {
	return cli.Command{
		Name:    "jobs",
		Aliases: []string{"j"},
		Usage:   "Open jobs page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project, p",
				Usage: "Specify the project to open",
			},
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
			cc, err := cc.GetProvider(org)
			if err != nil {
				return err
			}
			return browser.Open(c, cc)
		},
	}
}

func newCircleCIWorkflowsCmd() cli.Command {
	return cli.Command{
		Name:    "workflows",
		Aliases: []string{"wf"},
		Usage:   "Open workflows page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project, p",
				Usage: "Specify the project to open",
			},
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
			cc, err := cc.GetProvider(org)
			if err != nil {
				return err
			}
			return browser.Open(c, cc)
		},
	}
}
