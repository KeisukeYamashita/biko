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

package googleworkspace

import (
	"net/url"
	"path"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/urfave/cli"
)

const (
	drive        = "drive"
	document     = "document"
	spreadsheets = "spreadsheets"
	presentation = "presentation"
	forms        = "forms"
)

// Provider ...
type Provider struct {
	baseURL *url.URL
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
		Aliases: conf.GoogleWorkspace["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	newFlag := p.GetCtxString("new")
	if newFlag == "true" {  // HACK: need to fix GetCtxString
		return p.getNewCmdURL(), nil
	}

	var baseURL string
	product := p.Ctx.Command.Name
	switch product {
	case drive:
		baseURL = "https://drive.google.com"
	case document, spreadsheets, presentation, forms:
		baseURL = "https://docs.google.com/"
	}

	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}
	p.addProductPath(product)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	p.URL = p.baseURL
	switch product {
	case drive:
		p.join(drive)
		param := url.Values{}
		var query string
		if query = p.GetCtxString("query"); query != "" {
			p.join("search")
			param.Add("q", query)
			p.URL.RawQuery = param.Encode()
		}
	case document:
		p.join(document)
		param := url.Values{}
		var query string
		if query = p.GetCtxString("query"); query != "" {
			param.Add("q", query)
			p.URL.RawQuery = param.Encode()
		}
	case spreadsheets:
		p.join(spreadsheets)
		param := url.Values{}
		var query string
		if query = p.GetCtxString("query"); query != "" {
			param.Add("q", query)
			p.URL.RawQuery = param.Encode()
		}
	case presentation:
		p.join(presentation)
		param := url.Values{}
		var query string
		if query = p.GetCtxString("query"); query != "" {
			param.Add("q", query)
			p.URL.RawQuery = param.Encode()
		}
	case forms:
		p.join(forms)
		param := url.Values{}
		var query string
		if query = p.GetCtxString("query"); query != "" {
			param.Add("q", query)
			p.URL.RawQuery = param.Encode()
		}
	}
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}

func (p *Provider) getNewCmdURL() string {
	product := p.Ctx.Command.Name
	switch product {
	case document:
		return "https://document.new"
	case spreadsheets:
		return "https://spreadsheets.new"
	case presentation:
		return "https://presentation.new"
	case forms:
		return "https://forms.new"
	}
	return ""
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
