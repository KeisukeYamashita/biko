package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

const (
	version = "0.0.1"
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
