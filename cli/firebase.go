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
	"github.com/KeisukeYamashita/biko/providers/firebase"
	"github.com/urfave/cli"
)

const (
	categoryDevelop   = "Develop"
	categoryQuality   = "Quality"
	categoryAnalytics = "Analitics"
	categoryGrow      = "Grow"
)

func newFirebaseCmd() cli.Command {
	return cli.Command{
		Name:    "firebase",
		Aliases: []string{"fb"},
		Usage:   "Open Firebase source",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
		Subcommands: []cli.Command{
			newFirebaseAuthenticateCmd(),
			newFirebaseDatabaseCmd(),
			newFirebaseStorageCmd(),
			newFirebaseHostingCmd(),
			newFirebaseFunctionsCmd(),
			newFirebaseMLKitCmd(),
			newFirebaseCrashlyticsCmd(),
			newFirebasePerformanceCmd(),
			newFirebaseTestLabCmd(),
			newFirebaseAppDistributionCmd(),
			newFirebasePredictionsCmd(),
		},
	}
}

func newFirebaseAuthenticateCmd() cli.Command {
	return cli.Command{
		Name:     "authentication",
		Aliases:  []string{"auth"},
		Category: categoryDevelop,
		Usage:    "Open authentication page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseDatabaseCmd() cli.Command {
	return cli.Command{
		Name:     "database",
		Aliases:  []string{"db"},
		Category: categoryDevelop,
		Usage:    "Open database page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseStorageCmd() cli.Command {
	return cli.Command{
		Name:     "storage",
		Category: categoryDevelop,
		Usage:    "Open storage page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseHostingCmd() cli.Command {
	return cli.Command{
		Name:     "hosting",
		Aliases:  []string{"host", "h"},
		Category: categoryDevelop,
		Usage:    "Open storage page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseFunctionsCmd() cli.Command {
	return cli.Command{
		Name:     "functions",
		Aliases:  []string{"func", "f"},
		Category: categoryDevelop,
		Usage:    "Open functions page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseMLKitCmd() cli.Command {
	return cli.Command{
		Name:     "ml",
		Category: categoryDevelop,
		Usage:    "Open ML kit page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseCrashlyticsCmd() cli.Command {
	return cli.Command{
		Name:     "crashlytics",
		Aliases:  []string{"crash"},
		Category: categoryQuality,
		Usage:    "Open crashlytics page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebasePerformanceCmd() cli.Command {
	return cli.Command{
		Name:     "performance",
		Aliases:  []string{"perf"},
		Category: categoryQuality,
		Usage:    "Open performance page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseTestLabCmd() cli.Command {
	return cli.Command{
		Name:     "testlab",
		Aliases:  []string{"tl"},
		Category: categoryQuality,
		Usage:    "Open testlab page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseAppDistributionCmd() cli.Command {
	return cli.Command{
		Name:     "appdistribution",
		Aliases:  []string{"ad"},
		Category: categoryQuality,
		Usage:    "Open app distribution page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebaseDashboardCmd() cli.Command {
	return cli.Command{
		Name:     "dashboard",
		Aliases:  []string{"a"},
		Category: categoryAnalytics,
		Usage:    "Open dashboard page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}

func newFirebasePredictionsCmd() cli.Command {
	return cli.Command{
		Name:     "predictions",
		Aliases:  []string{"pred"},
		Category: categoryGrow,
		Usage:    "Open predictions page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			fb, err := firebase.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, fb)
		},
	}
}
