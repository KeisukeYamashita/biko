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
	dd "github.com/KeisukeYamashita/biko/providers/datadog"
	"github.com/urfave/cli"
)

func newDatadaogCmd() cli.Command {
	return cli.Command{
		Name:     "datadog",
		Aliases:  []string{"dd"},
		Usage:    "Open Datadog resource",
		Category: categoryMonitor,
		Flags:    []cli.Flag{},
		Subcommands: []cli.Command{
			newDDWatchDogCmd(),
			newDDEventCmd(),
			newDDDashboardCmd(),
			newDDInfrastructureCmd(),
			newDDMonitorsCmd(),
			newDDIntegrationsCmd(),
			newDDApmCmd(),
			newDDNotebookCmd(),
			newDDLogsCmd(),
			newDDSyntheticsCmd(),
		},
	}
}

func newDDWatchDogCmd() cli.Command {
	return cli.Command{
		Name:    "watchdog",
		Aliases: []string{"wd"},
		Usage:   "Open Watchdog page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDEventCmd() cli.Command {
	return cli.Command{
		Name:  "events",
		Usage: "Open Events page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDDashboardCmd() cli.Command {
	return cli.Command{
		Name:  "dashboard",
		Usage: "Open Dashboard page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDInfrastructureCmd() cli.Command {
	return cli.Command{
		Name:  "infrastructure",
		Usage: "Open Infrastructure page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDMonitorsCmd() cli.Command {
	return cli.Command{
		Name:  "monitors",
		Usage: "Open Monitors page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDIntegrationsCmd() cli.Command {
	return cli.Command{
		Name:  "integrations",
		Usage: "Open Integrations page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDApmCmd() cli.Command {
	return cli.Command{
		Name:  "apm",
		Usage: "Open APM page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDNotebookCmd() cli.Command {
	return cli.Command{
		Name:  "notebook",
		Usage: "Open Notebook page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDLogsCmd() cli.Command {
	return cli.Command{
		Name:  "logs",
		Usage: "Open Logs page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "view, v",
				Usage: "Specify the saved view to open",
			},
		},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}

func newDDSyntheticsCmd() cli.Command {
	return cli.Command{
		Name:  "synthetics",
		Usage: "Open Synthetics page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd, err := dd.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, dd)
		},
	}
}
