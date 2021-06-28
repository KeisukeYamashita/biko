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
	"github.com/KeisukeYamashita/biko/providers/gcp"
	"github.com/urfave/cli"
)

const (
	categoryCommon      = "Common"
	categoryCompute     = "Compute"
	categoryStorage     = "Storage"
	categoryNetworking  = "Networking"
	categoryStackdriver = "Stackdriver"
	categoryTools       = "Tools"
	categoryBigData     = "Big Data"
)

func newGCPCmd() cli.Command {
	return cli.Command{
		Name:     "gcp",
		Usage:    "Open GCP resource",
		Category: categoryCloudProvider,
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
			newGCPCloudRunCmd(),
			newGCPDataflowCmd(),
			newGCPFunctionsCmd(),
			newGCPGCECmd(),
			newGCPGCRCmd(),
			newGCPGKECmd(),
			newGCPIAMCmd(),
			newGCPLogsCmd(),
			newGCPKMSCmd(),
			newGCPPubSubCmd(),
			newGCPSecretManagerCmd(),
			newGCPSpannerCmd(),
			newGCPStorageCmd(),
			newGCPSQLCmd(),
		},
	}
}

func newGCPGAECmd() cli.Command {
	return cli.Command{
		Name:     "appengine",
		Aliases:  []string{"gae"},
		Category: categoryCompute,
		Usage:    "Open GAE page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPBQCmd() cli.Command {
	return cli.Command{
		Name:     "bigquery",
		Aliases:  []string{"bq"},
		Category: categoryBigData,
		Usage:    "Open Bigquery page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "database, db",
				Usage: "Database to show",
			},
			cli.StringFlag{
				Name:  "table, t",
				Usage: "Table to show",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGKECmd() cli.Command {
	return cli.Command{
		Name:     "kubernetes",
		Aliases:  []string{"gke"},
		Category: categoryCompute,
		Usage:    "Open GKE page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Region of the cluster",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name of the cluter",
			},
			cli.StringFlag{
				Name:  "namespaces, ns",
				Usage: "Namespaces of workload page to open (you can input multiple namespaces by comma-speparated string)",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPSpannerCmd() cli.Command {
	return cli.Command{
		Name:     "spanner",
		Usage:    "Open Cloud Spanner page",
		Category: categoryStorage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "instance, i",
				Usage: "Instance name",
			},
			cli.StringFlag{
				Name:  "database, db",
				Usage: "Database name",
			},
			cli.StringFlag{
				Name:  "table, tb",
				Usage: "Table name",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPSecretManagerCmd() cli.Command {
	return cli.Command{
		Name:     "secretManager",
		Usage:    "Open Secret Manager page",
		Category: categoryStorage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "secret, s",
				Usage: "Secret name",
			},
			cli.StringFlag{
				Name:  "version, v",
				Usage: "Secret version",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGCRCmd() cli.Command {
	return cli.Command{
		Name:     "gcr",
		Usage:    "Open Cloud Registry page",
		Category: categoryTools,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name of the image",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPFunctionsCmd() cli.Command {
	return cli.Command{
		Name:     "functions",
		Aliases:  []string{"f"},
		Category: categoryCompute,
		Usage:    "Open Cloud Functions page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name of the function",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Region of the function",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPCloudRunCmd() cli.Command {
	return cli.Command{
		Name:     "run",
		Category: categoryCompute,
		Usage:    "Open Cloud Run page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name of the function",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Region of the function",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPGCECmd() cli.Command {
	return cli.Command{
		Name:     "compute",
		Aliases:  []string{"gce"},
		Category: categoryCompute,
		Usage:    "Open Compute Engine page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPLogsCmd() cli.Command {
	return cli.Command{
		Name:     "logs",
		Aliases:  []string{"l"},
		Category: categoryStackdriver,
		Usage:    "Open Stackdriver log page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPIAMCmd() cli.Command {
	return cli.Command{
		Name:     "iam",
		Category: categoryCommon,
		Usage:    "Open IAM & admin page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPSQLCmd() cli.Command {
	return cli.Command{
		Name:     "sql",
		Category: categoryStorage,
		Usage:    "Open Cloud SQL page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPPubSubCmd() cli.Command {
	return cli.Command{
		Name:     "pubsub",
		Category: categoryBigData,
		Usage:    "Open Cloud PubSub page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPStorageCmd() cli.Command {
	return cli.Command{
		Name:     "storage",
		Aliases:  []string{"gcs"},
		Category: categoryStorage,
		Usage:    "Open Cloud Storage page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
			cli.StringFlag{
				Name:  "bucket, b",
				Usage: "Specify the bucket to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPDataflowCmd() cli.Command {
	return cli.Command{
		Name:     "dataflow",
		Category: categoryBigData,
		Usage:    "Open Cloud Dataflow page",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}

func newGCPKMSCmd() cli.Command {
	return cli.Command{
		Name:     "kms",
		Category: categoryBigData,
		Usage:    "Open Cloud KMS",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project",
				Usage: "Specify the project to open",
			},
		},
		Action: func(c *cli.Context) error {
			gcp, err := gcp.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, gcp)
		},
	}
}
