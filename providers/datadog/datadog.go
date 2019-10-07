package datadog

import (
	"net/url"
	"path"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/urfave/cli"
)

// Provider ...
type Provider struct {
	baseURL *url.URL
	URL     *url.URL
	Product string
	Ctx     *cli.Context
	Aliases map[string]interface{}
}

// GetProvider ...
func GetProvider() (*Provider, error) {
	conf, err := alias.GetConfig()
	if err != nil {
		return nil, err
	}

	return &Provider{
		Aliases: conf.Datadog["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	const baseURL = "https://app.datadoghq.com"
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}
func (p *Provider) addProductPath(product string) {
	p.join(product)
	switch product {
	case "watchdogs":
	case "events":
		p.join("events/stream")
	case "metrics":
	case "integrations":
		p.join("account/settings")
	case "apm":
	case "notebook":
	case "logs":
	case "synthetics":
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

// GetCtxString ...
func (p *Provider) GetCtxString(str string) string {
	key := p.GetCtxString(str)
	if key == "" {
		return ""
	}
	value, ok := p.Aliases[key].(string)
	if !ok {
		return key
	}
	if value == "" {
		return key
	}
	return value
}
