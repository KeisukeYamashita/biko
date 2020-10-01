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
	"net/url"

	"github.com/KeisukeYamashita/biko/browser"
	jr "github.com/KeisukeYamashita/biko/providers/jira"
	"github.com/urfave/cli"
)

func newJIRACmd() cli.Command {
	return cli.Command{
		Name:     "jira",
		Aliases:  []string{"jr"},
		Usage:    "Open JIRA resource",
		Category: categoryCloudProvider,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
		Subcommands: []cli.Command{
			newJIRADashboardCmd(),
			newJIRAProjectsCmd(),
			newJIRAPeopleCmd(),
			newJIRAIssuesCmd(),
			newJIRABacklogCmd(),
			newJIRAReportsCmd(),
		},
	}
}

func newJIRADashboardCmd() cli.Command {
	return cli.Command{
		Name:    "dashboard",
		Aliases: []string{"db"},
		Usage:   "Open dashboard page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func newJIRAProjectsCmd() cli.Command {
	return cli.Command{
		Name:    "projects",
		Aliases: []string{"ps"},
		Usage:   "Open projects page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func newJIRAPeopleCmd() cli.Command {
	return cli.Command{
		Name:    "people",
		Aliases: []string{"pp"},
		Usage:   "Open people page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func newJIRAIssuesCmd() cli.Command {
	return cli.Command{
		Name:    "issues",
		Aliases: []string{"is"},
		Usage:   "Open issues page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func newJIRABacklogCmd() cli.Command {
	return cli.Command{
		Name:    "backlog",
		Aliases: []string{"bl"},
		Usage:   "Open backlog page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "project, p",
				EnvVar: "BIKO_JIRA_PROJECT",
				Usage:  "Specify the project to open",
			},
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func newJIRAReportsCmd() cli.Command {
	return cli.Command{
		Name:    "reports",
		Aliases: []string{"rp"},
		Usage:   "Open reports page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "project, p",
				EnvVar: "BIKO_JIRA_PROJECT",
				Usage:  "Specify the project to open",
			},
			cli.StringFlag{
				Name:   "org",
				EnvVar: "BIKO_JIRA",
				Usage:  "Specify JIRA Organization",
			},
			cli.StringFlag{
				Name:   "base",
				EnvVar: "BIKO_JIRA_BASE",
				Usage:  "Specify JIRA Base URL (for self-managed plan)",
			},
		},
		Action: func(c *cli.Context) error {
			var project string
			if project = c.String("project"); project == "" {
				return fmt.Errorf("Project for jira not configured pass --project or set BIKO_JIRA_PROJECT")
			}
			baseURL, err := getBaseURL(c)
			if err != nil {
				return err
			}
			jr := &jr.Provider{}
			if jr.BaseURL, err = url.Parse(baseURL); err != nil {
				return err
			}
			return browser.Open(c, jr)
		},
	}
}

func getBaseURL(c *cli.Context) (string, error) {
	org := c.String("org")
	base := c.String("base")
	if org == "" && base == "" {
		return "", fmt.Errorf("Org and Base for jira not configured pass --org/BIKO_JIRA or --base/BIKO_JIRA_BASE")
	}
	var baseURL = fmt.Sprintf("https://%s.atlassian.net", org) // for jira cloud
	if base != "" {
		baseURL = base // for self-managed
	}
	return baseURL, nil
}
