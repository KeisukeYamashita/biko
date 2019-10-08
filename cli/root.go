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
	"os"

	"github.com/urfave/cli"
)

// Run creates command and run the command
func Run() error {
	cmd := NewCmdRoot()
	return cmd.Run(os.Args)
}

// NewCmdRoot will create root command
func NewCmdRoot() *cli.App {
	cmd := cli.NewApp()
	cmd.Version = version
	cmd.Commands = rootSubCommands()
	return cmd
}

func rootSubCommands() []cli.Command {
	return []cli.Command{
		newVersionCmd(),
		newGCPCmd(),
		newDatadaogCmd(),
		newGoogleCmd(),
		newYoutubeCmd(),
		newPagerDutyCmd(),
		newGithubCmd(),
		newCircleCICmd(),
	}
}
