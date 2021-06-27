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
	"github.com/KeisukeYamashita/biko/providers/googleworkspace"
	"github.com/urfave/cli"
)

func newGoogleWorkspaceCmd() cli.Command {
	return cli.Command{
		Name:     "googleworkspace",
		Aliases:  []string{"gw"},
		Usage:    "Open Google Workspace resource",
		Category: categoryWebService,
		Flags:    []cli.Flag{},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
		Subcommands: []cli.Command{
			newDriveCmd(),
			newDocumentCmd(),
			newSpreadsheetsCmd(),
			newPresentationCmd(),
			newFormsCmd(),
		},
	}
}

func newDriveCmd() cli.Command {
	return cli.Command{
		Name:    "drive",
		Aliases: []string{"dr"},
		Usage:   "Open Google Drive directory",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query to search",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}

func newDocumentCmd() cli.Command {
	return cli.Command{
		Name:    "document",
		Aliases: []string{"doc"},
		Usage:   "Open Google Document page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
			cli.BoolFlag{
				Name:  "new, n",
				Usage: "Create a new document (this flag prioritize over query flag)",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}

func newSpreadsheetsCmd() cli.Command {
	return cli.Command{
		Name:    "spreadsheets",
		Aliases: []string{"ss"},
		Usage:   "Open Google Spreadsheets page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
			cli.BoolFlag{
				Name:  "new, n",
				Usage: "Create a new spreadsheet (this flag prioritize over query flag)",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}

func newPresentationCmd() cli.Command {
	return cli.Command{
		Name:    "presentation",
		Aliases: []string{"pr"},
		Usage:   "Open Google Slides page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
			cli.BoolFlag{
				Name:  "new, n",
				Usage: "Create a new presentation (this flag prioritize over query flag)",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}

func newFormsCmd() cli.Command {
	return cli.Command{
		Name:    "forms",
		Aliases: []string{"fm"},
		Usage:   "Open Google Forms page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query, q",
				Usage: "Query a page",
			},
			cli.BoolFlag{
				Name:  "new, n",
				Usage: "Create a new form (this flag prioritize over query flag)",
			},
		},
		Action: func(c *cli.Context) error {
			g, err := googleworkspace.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, g)
		},
	}

}
