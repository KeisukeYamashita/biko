# Biko

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

Open provider.

```
$ biko gcp --project my-spanner-project --product spanner
```

If you don't specify the project, it will use the value configured by `gcloud` SDK config.

## Support

The supported provider are here.

* GoogleCloudPlatform

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