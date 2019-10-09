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
	az "github.com/KeisukeYamashita/biko/providers/azure"
	"github.com/urfave/cli"
)

const (
	categoryAZuComputing = "Computing"
)

func newAzureCmd() cli.Command {
	return cli.Command{
		Name:    "azure",
		Aliases: []string{"az"},
		Usage:   "Open Microsoft Azure resource",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			aws, err := az.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
		Subcommands: []cli.Command{},
	}
}
