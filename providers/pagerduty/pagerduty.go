package pagerduty

import (
	"fmt"
	"net/url"
	"path"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/urfave/cli"
)

// Provider ...
type Provider struct {
	baseURL *url.URL
	URL     *url.URL
	Ctx     *cli.Context
	Org     string
	Aliases map[string]interface{}
}

// GetProvider ...
func GetProvider() (*Provider, error) {
	conf, err := alias.GetConfig()
	if err != nil {
		return nil, err
	}

	return &Provider{
		Aliases: conf.PagerDuty["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	var baseURL = fmt.Sprintf("https://%s.pagerduty.com", p.Org)
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	switch product {
	case "incidents":
		p.join("incidents")
	case "alerts":
		p.join("alerts")
	case "schedules":
		p.join("schedules")
	default:
		p.join("incidents")
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
