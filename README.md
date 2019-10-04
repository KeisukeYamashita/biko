# Biko

![CircleCI](https://circleci.com/gh/KeisukeYamashita/biko.svg?style=svg&circle-token=e4b3002b1fb96c423ed5f75332c3455d88d56b0f)
![License](https://img.shields.io/badge/license-Apache%202.0-%23E93424)

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
* Youtube
* Pagerduty
* Github

### GCP

TODO:

### Datadog

TODO:

### Google

TODO:

### Youtube

TODO: 

### Pagerduty

TODO:

### Github

TODO: 

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