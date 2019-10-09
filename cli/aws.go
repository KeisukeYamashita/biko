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
	"github.com/KeisukeYamashita/biko/providers/aws"
	"github.com/urfave/cli"
)

const (
	categoryAWSComputing  = "Computing"
	categoryAWSStorage    = "Storage"
	categoryAWSDatabase   = "Database"
	categoryAWSNetworking = "Networking and Contens Delivery"
)

func newAWSCmd() cli.Command {
	return cli.Command{
		Name:  "aws",
		Usage: "Open AWS resource",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
		Subcommands: []cli.Command{
			newAWSEC2Cmd(),
			newAWSECRCmd(),
			newAWSECSCmd(),
			newAWSLambdaCmd(),
			newAWSBatchCmd(),
			newAWSS3Cmd(),
			newAWSEFSCmd(),
			newAWSRDSCmd(),
			newAWSDynamoDBCmd(),
			newAWSNeptuneCmd(),
			newAWSVPCCmd(),
			newAWSRoute53Cmd(),
		},
	}
}

func newAWSEC2Cmd() cli.Command {
	return cli.Command{
		Name:     "ec2",
		Usage:    "Open EC2 page",
		Category: categoryAWSComputing,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSECRCmd() cli.Command {
	return cli.Command{
		Name:     "ecr",
		Usage:    "Open ECR page",
		Category: categoryAWSComputing,
		Flags:    []cli.Flag{},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSECSCmd() cli.Command {
	return cli.Command{
		Name:     "ecs",
		Usage:    "Open ECS page",
		Category: categoryAWSComputing,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSEKSCmd() cli.Command {
	return cli.Command{
		Name:     "eks",
		Usage:    "Open EKS page",
		Category: categoryAWSComputing,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSLambdaCmd() cli.Command {
	return cli.Command{
		Name:     "lambda",
		Aliases:  []string{"lam", "l"},
		Usage:    "Open EKS page",
		Category: categoryAWSComputing,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSBatchCmd() cli.Command {
	return cli.Command{
		Name:     "batch",
		Aliases:  []string{"b"},
		Usage:    "Open Batch page",
		Category: categoryAWSComputing,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSS3Cmd() cli.Command {
	return cli.Command{
		Name:     "s3",
		Usage:    "Open S3 page",
		Category: categoryAWSStorage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSEFSCmd() cli.Command {
	return cli.Command{
		Name:     "efs",
		Usage:    "Open EFS page",
		Category: categoryAWSStorage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSRDSCmd() cli.Command {
	return cli.Command{
		Name:     "rds",
		Usage:    "Open RDS page",
		Category: categoryAWSDatabase,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSDynamoDBCmd() cli.Command {
	return cli.Command{
		Name:     "dynamodb",
		Aliases:  []string{"dynamo", "dyn"},
		Usage:    "Open Dynamo page",
		Category: categoryAWSDatabase,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSNeptuneCmd() cli.Command {
	return cli.Command{
		Name:     "neptune",
		Usage:    "Open Neptune page",
		Category: categoryAWSDatabase,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSRedshiftCmd() cli.Command {
	return cli.Command{
		Name:     "redshift",
		Aliases:  []string{"red", "rs"},
		Usage:    "Open Neptune page",
		Category: categoryAWSDatabase,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSVPCCmd() cli.Command {
	return cli.Command{
		Name:     "vpc",
		Usage:    "Open VPC page",
		Category: categoryAWSNetworking,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}

func newAWSRoute53Cmd() cli.Command {
	return cli.Command{
		Name:     "route53",
		Usage:    "Open Route53 page",
		Aliases:  []string{"route", "53"},
		Category: categoryAWSNetworking,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Specify the region to open",
			},
		},
		Action: func(c *cli.Context) error {
			aws, err := aws.GetProvider()
			if err != nil {
				return err
			}
			return browser.Open(c, aws)
		},
	}
}
