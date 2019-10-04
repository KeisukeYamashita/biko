package providers

import "github.com/urfave/cli"

// Provider are interfaces ...
type Provider interface {
	Init(c *cli.Context) error
	GetTargetURL() (string, error)
}
