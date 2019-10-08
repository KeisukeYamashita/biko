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
	"github.com/KeisukeYamashita/biko/providers/google"
	"github.com/urfave/cli"
)

func newGoogleCmd() cli.Command {
	return cli.Command{
		Name:    "google",
		Aliases: []string{"g"},
		Usage:   "Open Google source",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := google.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
		Subcommands: []cli.Command{
			newGoogleSearchCmd(),
		},
	}
}

func newGoogleSearchCmd() cli.Command {
	return cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "Search a page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := google.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}
