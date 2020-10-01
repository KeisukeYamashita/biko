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

package jira

import (
	"fmt"
	"net/url"
	"path"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/urfave/cli"
)

// Provider ...
type Provider struct {
	BaseURL *url.URL
	URL     *url.URL
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
		Aliases: conf.JIRA["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	switch product {
	case "dashboard":
		p.join("secure/Dashboard.jspa")
	case "projects":
		p.join("secure/BrowseProjects.jspa")
	case "people":
		p.join(fmt.Sprintf("jira/%s", product))
	case "issues":
		p.join(product)
	case "backlog":
		p.join("browse")
		if project := p.GetCtxString("project"); project != "" {
			p.join(project)
		}
	case "reports":
		project := p.GetCtxString("project")
		p.join(fmt.Sprintf("projects/%s?selectedItem=com.atlassian.jira.jira-projects-plugin%%3Areport-page", project))
	default:
		p.join("secure/Dashboard.jspa")
	}

	return
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.BaseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
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
