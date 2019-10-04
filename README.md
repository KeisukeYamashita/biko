# Biko

![CircleCI](https://circleci.com/gh/KeisukeYamashita/biko.svg?style=svg&circle-token=e4b3002b1fb96c423ed5f75332c3455d88d56b0f)

<br />
<p align="center"><a href="#" target="_blank" rel="noopener noreferrer"><img width="320"src="_image/logo.png"></a></p>
<br />

## Install

### By Homebrew

Install by homebrew.

```
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
$ biko gcp --project my-spanner-project --product spanner
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

### GCP

TODO:

### Datadog

TODO:

### Google

TODO:

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

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)