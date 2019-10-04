package github

import (
	"fmt"
	"net/url"
	"path"

	"github.com/urfave/cli"
)

type Provider struct {
	baseURL *url.URL
	URL     *url.URL
	Ctx     *cli.Context
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	const baseURL = "https://github.com"
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	switch product {
	case "dashboard":
		var org string
		if org = p.Ctx.String("organization"); org != "" {
			p.join(fmt.Sprintf("orgs/%s/dashboard", org))
			return
		}
		p.URL = p.baseURL
		return
	default:
		p.URL = p.baseURL
	}
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}
