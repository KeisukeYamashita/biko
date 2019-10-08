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
		Usage:   "Open the trending page",
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
