package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	dd "github.com/KeisukeYamashita/biko/providers/datadog"
	"github.com/urfave/cli"
)

func newDatadaogCmd() cli.Command {
	return cli.Command{
		Name:    "datadog",
		Aliases: []string{"dd"},
		Usage:   "Open Datadog resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "product",
				Usage: "Product to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &dd.Provider{
				Product: c.String("product"),
			}
			return browser.Open(c, gcp)
		},
	}
}
