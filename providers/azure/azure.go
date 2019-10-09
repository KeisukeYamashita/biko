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

package azure

import (
	"fmt"
	"path"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/urfave/cli"
)

// Provider ...
// TODO(KeisukeYamashita): Use url package for baseURL and URL
type Provider struct {
	baseURL string
	URL     string
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
		Aliases: conf.Azure["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	p.baseURL = "https://portal.azure.com"
	p.addProductPath(p.Ctx.Command.Name)
	return p.URL, nil
}

func (p *Provider) addProductPath(product string) {
	p.URL = p.baseURL
	if product == "" {
		p.join("#home")
		return
	}

	var productType, path string
	suffix := "Blade"
	switch product {
	case "vm":
		productType = "Compute"
		path = "VirtualMachines"
	case "appservices":
		productType = "Web"
		path = "sites"
	case "funcion":
		productType = "Web"
		path = "sites/king/functionapp"
	case "sql":
		productType = "Sql"
		path = "services/databases"
	case "cosmos":
		productType = "DocumentDB"
		path = "databaseAccounts"
		suffix = ""
	case "storage":
		productType = "Storage"
		path = "StorageAccounts"
	default:
		return
	}
	p.join(fmt.Sprintf("#blade/HubsExtension/BrowseResource%s/resourceType/Microsoft.%s%%2F%s", suffix, productType, path))
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
	if p.URL == "" {
		p.URL = p.baseURL
	}
	p.URL = path.Join(p.URL, additionPath)
}
