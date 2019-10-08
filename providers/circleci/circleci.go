// Copyright 2019 The Biko Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
