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

package gcp

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/KeisukeYamashita/biko/alias"
	"github.com/go-ini/ini"
	"github.com/urfave/cli"
)

const (
	baseURL = "https://console.cloud.google.com"
)

var (
	defaultGoogleSDKConfigPath = os.Getenv("HOME") + "/.config/gcloud/configurations/config_default"
)

// Provider ...
type Provider struct {
	baseURL   *url.URL
	URL       *url.URL
	SDKConfig *SDKConfig
	Product   string
	Ctx       *cli.Context
	Aliases   map[string]interface{}
}

// SDKConfig ...
type SDKConfig struct {
	Core    *Core
	Compute *Compute
	Cluster *Container
}

// Core ...
type Core struct {
	Account string
	Project string
}

// Compute ...
type Compute struct {
	Zone   string
	Region string
}

// Container ...
type Container struct {
	Cluster string
}

// GetProvider ...
func GetProvider() (*Provider, error) {
	conf, err := alias.GetConfig()
	if err != nil {
		return nil, err
	}

	return &Provider{
		Aliases: conf.GCP["alias"].(map[string]interface{}),
	}, nil
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	var err error
	if p.SDKConfig, err = getSDKConfig(); err != nil {
		return err
	}

	if c.String("project") != "" {
		p.SDKConfig.Core.Project = c.String("project")
	}

	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Ctx.Command.Name)
	p.addProjectParam()
	return p.URL.String(), nil
}

func getSDKConfig() (*SDKConfig, error) {
	cfg, err := ini.Load(defaultGoogleSDKConfigPath)
	if err != nil {
		return nil, err
	}

	conf := &SDKConfig{
		&Core{
			Account: cfg.Section("core").Key("account").MustString("xxx@email.com"),
			Project: cfg.Section("core").Key("project").MustString("xxx"),
		},
		&Compute{
			Zone:   cfg.Section("compute").Key("zone").MustString("xxx"),
			Region: cfg.Section("compute").Key("region").MustString("xxx"),
		},
		&Container{
			Cluster: cfg.Section("container").Key("cluster").MustString("xxx"),
		},
	}

	return conf, nil
}

func (p *Provider) addProductPath(product string) {
	switch product {
	case "appengine":
		p.join(product)
	case "bigquery":
		p.join(product)
		var db, table string
		if db = p.GetCtxString("database"); db != "" {
			q := p.URL.Query()
			q.Add("d", db)
			if table = p.GetCtxString("table"); table != "" {
				q.Add("t", table)
			}
			p.URL.RawQuery = q.Encode()
		}
	case "kubernetes":
		p.join(product)
		var region, name, namespaces string
		if region = p.GetCtxString("region"); region != "" {
			p.join(fmt.Sprintf("details/%s", region))
			if name = p.GetCtxString("name"); name != "" {
				p.join(name)
			}
		} else if namespaces = p.GetCtxString("namespaces"); namespaces != "" {
			p.join("workload")
			p.addGKEPageStateParam(namespaces)
		}
	case "secret-manager":
		p.join(fmt.Sprintf("security/%s", product))
		var secret string
		if secret = p.GetCtxString("secret"); secret != "" {
			p.join(fmt.Sprintf("secret/%s", secret))
		}
	case "spanner":
		p.join(product)
		var instance, db, scheme string
		if instance = p.GetCtxString("instance"); instance != "" {
			p.join(fmt.Sprintf("instances/%s", instance))
			if db = p.GetCtxString("database"); db != "" {
				p.join(fmt.Sprintf("databases/%s", db))
				if scheme = p.GetCtxString("table"); scheme != "" {
					p.join(fmt.Sprintf("schema/%s", scheme))
				}
			}
		}
	case "gcr":
		p.join(product)
		var name string
		p.join(fmt.Sprintf("images/%s/", p.SDKConfig.Core.Project))
		if name = p.GetCtxString("name"); name != "" {
			p.join(fmt.Sprintf("GLOBAL/%s", name))
		}
	case "run", "functions":
		p.join(product)
		var region, name string
		if region = p.GetCtxString("region"); region != "" {
			p.join(fmt.Sprintf("details/%s", region))

			if name = p.GetCtxString("name"); name != "" {
				p.join(name)
			}
		}

		switch product {
		case "run":
		case "functions":
		}
	case "logs":
		p.join(product)
	case "iam":
		p.join("iam-admin")
	case "sql":
		p.join(product)
	case "pubsub":
		p.join("cloudpubsub")
	case "storage":
		p.join(fmt.Sprintf("%s/browser", product))
		var bucket string
		if bucket = p.GetCtxString("bucket"); bucket != "" {
			p.join(bucket)
		}
	case "dataflow":
		p.join(product)
	case "kms":
		p.join("security/kms")
	default:
		p.join("home/dashboard")
	}

	return
}

func (p *Provider) addProjectParam() {
	q := p.URL.Query()
	q.Add("project", p.SDKConfig.Core.Project)
	p.URL.RawQuery = q.Encode()
	return
}

func (p *Provider) addGKEPageStateParam(namespaces string) {
	q := p.URL.Query()
	q.Add("pageState", constructPageStateParam(namespaces))
	p.URL.RawQuery = q.Encode()
	return
}

func constructPageStateParam(rawNamespaces string) string {
	nss := strings.Split(rawNamespaces, ",")
	quotedNamespaces := make([]string, 0, len(nss))
	for _, s := range nss {
		quotedNamespaces = append(quotedNamespaces, strconv.Quote(s))
	}
	return fmt.Sprintf("(\"savedViews\":(\"n\":[%s]))", strings.Join(quotedNamespaces, ","))
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
