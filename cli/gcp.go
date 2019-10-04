package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	"github.com/KeisukeYamashita/biko/providers/gcp"
	"github.com/urfave/cli"
)

func newGCPCmd() cli.Command {
	return cli.Command{
		Name:  "gcp",
		Usage: "Open GCP resource",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "product, prod",
				Usage: "Specify the product to open",
			},
			cli.StringFlag{
				Name:  "project, proj",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: c.String("product"),
			}
			return browser.Open(c, gcp)
		},
	}
}
