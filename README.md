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
$ go get github.com/KeisukeYamashita/biko
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

## Support

The supported provider are here.

* GoogleCloudPlatform(GCP)
* Datadog
* Google
* Youtube
* Pagerduty
* Github

### GCP

* By default, it will open the project configured by `gcloud` command.

```
$ biko gcp [product] [flag(s)]
```

**Supported Product**

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| App Engine | Go to GAE Dashboard | `appengine`, `gae` | - |
| Bigquery | Go to Bigquery top or the database, table | `bigquery`, `bq` | `--database`, `--table` |
| Kubernetes | Go to GKE page, or the cluster detail | `kubernetes`, `gke` | `--region`, `--name` | 
| Spanner | Go to spanner page, or the instance, database, table | `spanner` | `--instance`, `--database`, `--table` | 
| Container Registry | Go to container registry or the container detail | `gcr` | `--name` |
| Cloud Functions | Go to Cloud Functions page or the functions detail | `functions`, `f` |  `--region`, `--name` |
| Cloud Run | Go to Cloud Run page or the deployments detail | `run` | `--region`, `--name` |	
| Compute Engine | Go to GCE page | `compute` | -  |
| Stackdriver Logging | Go to Stackdriver logging | `logs`, `l` | - | 

Note that there is `--project` command flag for all commands.

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
| Logs | Go to logs page | `logs` | - | 
| Synthetics | Go to synthetics page | `synthetics` | - |

### Google

* Open Google and search from your terminal.

```
$ biko google [product] [flag(s)]
# or
$ biko g [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Search | Search on Google | `search`, `s` | - |


If you bump into something you want to lookup when you are using the terminal...

```
$ biko g s "How to configure biko"
```

Blazing fast.

### Youtube

* Open Youtube and search from your terminal.

```
$ biko youtube [product] [flag(s)]
# or
$ biko yt [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Search | Search on Youtube | `search`, `s` | - |

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

### Github

* Open Github from your terminal.

```
$ biko github [product] [flag(s)]
# or
$ biko gh [product] [flag(s)]
```

| Product | What | Command | Flags(Optional) | 
|:----:|:----:|:----:|:----:|
| Dashboard | Open github Dashboard | `dashboard` | `--org` |

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

## Roadmap

* See `docs/ROADMAP.md` for the project's roadmap.
    * Proposal are welcome.

## License

Copyright 2019 The Biko Authors.
Biko is released under the Apache License 2.0.

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

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