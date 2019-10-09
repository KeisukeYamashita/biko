# Biko

> CLI tool to jump to your browser directly to improve productivity

<br />
<p align="center"><a href="#" target="_blank" rel="noopener noreferrer"><img width="320"src="_image/logo.png"></a></p>
<br />

[![Go][go-badge]][go]
![CircleCI](https://circleci.com/gh/KeisukeYamashita/biko.svg?style=svg&circle-token=e4b3002b1fb96c423ed5f75332c3455d88d56b0f)
[![GolangCI][golangci-badge]][golangci]
[![Go Report Card][go-report-card-badge]][go-report-card]
[![GoDoc][godoc-badge]][godoc]
[![Dependabot][dependabot-badge]][dependabot]
![License](https://img.shields.io/badge/license-Apache%202.0-%23E93424)


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Install](#install)
  - [By Homebrew](#by-homebrew)
  - [By source code](#by-source-code)
- [Usage](#usage)
  - [Go to your target directly](#go-to-your-target-directly)
    - [Exaple1. Open GCP Cloud Spanner to specified project](#exaple1-open-gcp-cloud-spanner-to-specified-project)
    - [Example2. Open Datadog Dashboad](#example2-open-datadog-dashboad)
- [Alias](#alias)
- [Support](#support)
  - [AWS](#aws)
  - [Azure](#azure)
  - [CircleCI](#circleci)
  - [Datadog](#datadog)
  - [Firebase](#firebase)
  - [GCP](#gcp)
  - [Github](#github)
  - [Google](#google)
  - [Pagerduty](#pagerduty)
  - [Youtube](#youtube)
- [(Advanced): Docker image](#advanced-docker-image)
- [Contribute](#contribute)
  - [Add a provider](#add-a-provider)
- [Maintainers](#maintainers)
- [Contributers](#contributers)
- [Roadmap](#roadmap)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Install

### By Homebrew

Install by homebrew.

```
$ brew tap KeisukeYamashita/biko
$ brew install biko
```

### By source code

Install by source code.

```
$ go get github.com/KeisukeYamashita/biko/cmd/biko
```

Then build it.

```
$ make install
```

## Usage 

### Go to your target directly

#### Exaple1. Open GCP Cloud Spanner to specified project

```
$ biko gcp spanner --project my-spanner-project
```

If you don't specify the project, it will use the value configured by `gcloud` SDK config.

#### Example2. Open Datadog Dashboad 

```
$ biko dd --dashboard DASHBOARD_ID
```

## Alias 

You can configure your alias for whatever you want for frequently used projects, products, names, etc. 
For example, if you use the Cloud Function `my-cloud-functions` in region `asia-northeast1`, you can configure it like below in `$HOME/.biko/config.toml`.

```toml
[gcp]
	[gcp.alias]
	mcf = "my-cloud-functions"
	an1 = "asia-northeast1"
```

Now, you can open your page like below using the configured alias.

```
$ biko gcp functions -r as1 -n mcf
```

You can directly go to your page without waiting to load pages many times.

## Support

The supported provider are here.

* Amazon Web Service(AWS)
* Microsoft Azure
* CircleCI
* Datadog
* Firebase
* GoogleCloudPlatform(GCP)
* Github
* Google
* Pagerduty
* Youtube

### AWS

* Open AWS Console product page
* To open region page, please pass `--region, -r` flag.

```
$ biko aws [product] [flag(s)]
```

**Supported Product**

**Computing**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| EC2 | Open EC2 | `ec2` | - |
| ECR | Open ECR | `ecr` | - |
| ECS | Open ECR | `ecs` | - |
| EKS | Open EKR | `eks` | - |
| Lambda | Open Lambda | `lambda, lam` | - |
| Batch | Open Batch | `batch, b` | - |

**Storage**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| S3 | Open S3 page | `s3` | - |
| EFS | Open EFS page | `efs` | - |

**Database**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| RDS | Open RDS | `rds` | - |
| DynamoDB | Open DynamoDB |`dynamo, dyn` | - |
| Neptune | Open Neptune | `neptune, nep` | - |
| Redshift | Open Redshift | `redshift, red, rs` | - |

**Networking**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| VPC | Open VPC | `vpc` | - |
| Route53 | Open Route53 |`route53, route, 53` | - |

### Azure

* Open Microsoft Azure console

```
$ biko azure [product] [flag(s)]
# or
$ biko az [product] [flag(s)]
```

**Supported Product**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Virtual Machines | Open VM | `vm` | - |
| App Services | Open App Services | `appservices`, `as`, `sites` | - |
| Function App | Open Function App | `functions, f` | - | 
| Storage Accounts | Open Storage Accounts | `storage` | - | 
| SQL Databases | Open SQL Database | `sql` | - |
| Cosmos DB | Open Cosmos Database | `cosmos` | - |

### CircleCI

* Open CircleCI from your terminal.
* You need to pass `--org` or configure `BIKO_CIRCLECI`

```
$ biko circleci [product] [flag(s)]
# or
$ biko cc [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Jobs | Open CircleCI Jobs | `jobs, j` | `--project`, `-p` |
| Workflows | Open CircleCI Workflows | `workflows, wf` | `--project`, `-p` |

### Datadog

* `datadog` or `dd` command is supported for Datadog.

```
$ biko datadog [product] [flag(s)]
# or
$ biko dd [product] [flag(s)]
```

**Supported Product**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Watchdog | Go to Watchdog Dashboard | `watchdog`, `wd` | - |
| Events | Go to Events Dashboard | `events` | - |
| Dashboard | Go to Dashboard page | `dashboard` | - |
| Infrastructure | Go to Infrastructure page | `infrastructure` | - | 
| Monitors | Go to Monitors Dashboard | `moniters` | - |
| Integrations | Go to Integrations page | `integrations` | - |
| APM | Go to APM page | `apm` | - |
| Notebook | Go to Notebook page | `notebook` | - | 
| Logs | Go to logs page | `logs` | `--view, -v` |
| Synthetics | Go to synthetics page | `synthetics` | - |


### Firebase

* Open Google Firebase form your terminal.
* Please pass `--project` to open supported products page.

```
$ biko firebase [product] [flag(s)]
# or
$ biko fb [product] [flag(s)]
```

**Supported Product**

**Development**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Authentication | Go to auth | `authentications, auth` | - |
| Database | Go to database | `database, db` | - | 
| Storage | Go to storage | `storage` | - |
| Hosting | Go to Hosting | `hosting, host, h` | - |
| Functions | Go to functions | `funcitons, func, f` | - |
| ML Kit | Go to ML Kit | `ml` | - |

**Quality**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Crashlytics | Go to cryshlytics | `crashlytics, crash` | - |
| Performance | Go to performance | `performance, perf` | - |
| Test Lab | Go to Test Lab | `testlab, tl` | - |
| App Distribution | Go to App Distribution | `appdistribution, ad` | - |

**Analytics**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Dashboard | Go to dashboard | `dashboard` | - |

**Grow**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Grow | Go to grow | `grow` | - |

### GCP

* By default, it will open the project configured by `gcloud` command.

```
$ biko gcp [product] [flag(s)]
```

**Supported Product**

**Common**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| IAM & admin | Go to IAM & admin logging | `iam` | - |

**Compute**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| App Engine | Go to GAE Dashboard | `appengine`, `gae` | - |
| Compute Engine | Go to GCE page | `compute` | -  |
| Kubernetes | Go to GKE page, or the cluster detail | `kubernetes`, `gke` | `--region, -r`, `--name, -n` | 
| Cloud Functions | Go to Cloud Functions page or the functions detail | `functions`, `f` |  `--region, -r`, `--name, -n` |
| Cloud Run | Go to Cloud Run page or the deployments detail | `run` | `--region, -r`, `--name, -n` |	

**Storage**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Cloud SQL | Go to Cloud SQL | `sql` | - |
| Cloud Storage | Go to Cloud Storage `storage` | - |
| Spanner | Go to spanner page, or the instance, database, table | `spanner` | `--instance, -i`, `--database, -db`, `--table, -t` | 

**Networking**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|

**Stackdriver**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Stackdriver Logging | Go to Stackdriver logging | `logs`, `l` | - | 

**Tools**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Container Registry | Go to container registry or the container detail | `gcr` | `--name, -n` |


**Big Data**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Bigquery | Go to Bigquery top or the database, table | `bigquery`, `bq` | `--database, -db`, `--table, -tb` |
| Cloud PubSub | Go to Cloud PubSub | `pubsub` | - |
| Cloud Dataflow | Go to Cloud Dataflow | `dataflow` | - |

Note that there is `--project` command flag for all commands.

### Github

* Open Github from your terminal.

```
$ biko github [product] [flag(s)]
# or
$ biko gh [product] [flag(s)]
```

| Product   | What                  | Command     | Flags(Optional) |
| :----:    | :----:                | :----:      | :----:          |
| Dashboard | Open github Dashboard | `dashboard`, `db` | `--org`         |
| Trending  | Open github Treinding | `trending`, `t`   | `--language, -l`, `--since, -s` |

### Google

* Open Google and search from your terminal.

```
$ biko google [product] [flag(s)]
# or
$ biko g [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Search | Search on Google | `search`, `s` | `--query, -q` |


If you bump into something you want to lookup when you are using the terminal...

```
$ biko g s -q "How to configure biko"
```

Blazing fast.

### Pagerduty

* If you are using SSO, you need to pass `--org` or configure `BIKO_PAGERDUTY`

```
$ biko pagerduty [product] [flag(s)]
# or
$ biko pd [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Incident | Go to incident page | `incident`, `i` | - |
| Alert | Go to alert page | `alert`, `a` | - |
| Schedules | Go to schedules page | `schedules`, `s` | - |

### Youtube

* Open Youtube and search from your terminal.

```
$ biko youtube [product] [flag(s)]
# or
$ biko yt [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Search | Search on Youtube | `search`, `s` | `--query, -q` |

## (Advanced): Docker image

You can execute the cli with docker image.

```
$ docker run biko help
```

## Contribute

### Add a provider

You can add a provider by implementing this interface.

```go
// Provider are interfaces ...
type Provider interface {
	Init(c *cli.Context) error
	GetTargetURL() (string, error)
}
```

* `Init` will initialize your provider struct
* `GetTargetURL` should return a URL which will the browser will open

## Maintainers

<div style="display:flex;flex-direction:row;align-items: flex-start">
    <img href="https://github.com/KeisukeYamashita" src="./_image/maintainers/keke.png" width="100" style="margin:10px">
    <img href="https://github.com/sapuri" src="./_image/maintainers/sapuri.png" width="100" style="margin:10px">
</div>

## Contributers

<div style="display:flex;flex-direction:row;align-items: flex-start">
    <img href="https://github.com/KeisukeYamashita" src="./_image/contributors/keke.png" width="100" style="margin:10px">
    <img href="https://github.com/sapuri" src="./_image/contributors/sapuri.png" width="100" style="margin:10px">
    <img href="https://github.com/shopetan" src="./_image/contributors/shopetan.png" width="100" style="margin:10px">
    <img href="https://github.com/micnncim " src="./_image/contributors/micnncim.png" width="100" style="margin:10px">
</div>

## Roadmap

* See `docs/ROADMAP.md` for the project's roadmap.
    * Proposal are welcome.

## License

Copyright 2019 The Biko Authors.
Biko is released under the Apache License 2.0.

<!-- badge links -->
[go]: https://golang.org/dl
[go-badge]: https://img.shields.io/badge/Go-1.13-blue

[godoc]: https://godoc.org/github.com/KeisukeYamashita/biko
[godoc-badge]: https://img.shields.io/badge/godoc.org-reference-blue.svg

[golangci]: https://golangci.com/r/github.com/KeisukeYamashita/biko
[golangci-badge]: https://golangci.com/badges/github.com/KeisukeYamashita/biko.svg

[go-report-card]: https://goreportcard.com/report/github.com/KeisukeYamashita/biko
[go-report-card-badge]: https://goreportcard.com/badge/github.com/KeisukeYamashita/biko

[dependabot]: https://dependabot.com 
[dependabot-badge]: https://badgen.net/badge/icon/Dependabot?icon=dependabot&label&color=blue
