package circleci

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
func GetProvider(org string) (*Provider, error) {
	conf, err := alias.GetConfig()
	if err != nil {
		return nil, err
	}

	return &Provider{
		Aliases: conf.CircleCI["alias"].(map[string]interface{}),
		Org:     org,
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	const baseURL = "https://circleci.com/"
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}
	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	p.URL = p.baseURL
	switch product {
	case "jobs":
		p.join(fmt.Sprintf("gh/%s", p.Org))
		var project string
		if project = p.GetCtxString("project"); project != "" {
			p.join(fmt.Sprintf("gh/%s/%s/", p.Org, project))
			return
		}
		return
	case "workflows":
		p.join(fmt.Sprintf("gh/%s/workflows", p.Org))
		var project string
		if project = p.GetCtxString("project"); project != "" {
			p.join(project)
			return
		}
		return
	default:
		return
	}
}

// GetCtxString ...
func (p *Provider) GetCtxString(str string) string {
	key := p.Ctx.String(str)
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

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}
