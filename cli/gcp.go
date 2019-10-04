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
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
		Subcommands: []cli.Command{
			newGCPGAECmd(),
			newGCPBQCmd(),
			newGCPGKECmd(),
			newGCPSpannerCmd(),
			newGCPGCRCmd(),
			newGCPFunctionsCmd(),
		},
	}
}

func newGCPGAECmd() cli.Command {
	return cli.Command{
		Name:  "gae",
		Usage: "Open GAE page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "appengine",
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPBQCmd() cli.Command {
	return cli.Command{
		Name:  "bq",
		Usage: "Open Bigquery page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "bq",
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGKECmd() cli.Command {
	return cli.Command{
		Name:  "gke",
		Usage: "Open GKE page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "gke",
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPSpannerCmd() cli.Command {
	return cli.Command{
		Name:  "spanner",
		Usage: "Open Cloud Spanner page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "instance",
				Usage: "Instance name",
			},
			cli.StringFlag{
				Name:  "database",
				Usage: "Database name",
			},
		},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "spanner",
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGCRCmd() cli.Command {
	return cli.Command{
		Name:  "gcr",
		Usage: "Open Cloud Registry page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "gcr",
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPFunctionsCmd() cli.Command {
	return cli.Command{
		Name:  "functions",
		Usage: "Open Cloud Functions page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{
				Product: "cloudfunctions",
			}
			return browser.Open(c, gcp)
		},
	}
}
