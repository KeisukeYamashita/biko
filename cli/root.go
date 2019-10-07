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
