package gcp

import (
	"net/url"
	"path"

	"github.com/urfave/cli"
)

// Provider ...
type Provider struct {
	baseURL *url.URL
	URL     *url.URL
	Product string
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	const baseURL = "https://app.datadoghq.com"
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Product)
	return p.URL.String(), nil
}
func (p *Provider) addProductPath(product string) {
	switch product {
	case "watchdogs", "wd":
		p.join("watchdog")
	case "events":
		p.join("events/stream")
	case "dashboard", "dash", "d":
		p.join("dashboard")
	case "infrastructure", "infra":
		p.join("infrastructure")
	case "monitors", "monitor":
		p.join("monitors")
	case "metrics":
		p.join("metrics")
	case "integrations", "inte":
		p.join("account/settings")
	case "apm":
		p.join("apm")
	case "notebook", "nb":
		p.join("notebook")
	case "logs", "log":
		p.join("logs")
	case "synthetics":
		p.join("synthetics")
	default:
		p.join("apm/home")
	}
	return
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}
