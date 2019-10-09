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

	"github.com/urfave/cli"
)

const (
	version = "0.0.13"
)

func newVersionCmd() cli.Command {
	return cli.Command{
		Name:    "version",
		Usage:   "Show biko version",
		Aliases: []string{"v"},
		Action: func(c *cli.Context) error {
			fmt.Printf("Biko v%s\n", version)
			return nil
		},
	}
}
