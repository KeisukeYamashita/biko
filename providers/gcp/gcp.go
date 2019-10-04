package gcp

import (
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/go-ini/ini"
	"github.com/urfave/cli"
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
	const baseURL = "https://console.cloud.google.com"
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
	p.join(product)
	switch product {
	case "appengine":
	case "bigquery":
	case "kubernetes":
	case "spanner":
		var instance, db, scheme string
		if instance = p.Ctx.String("instance"); instance != "" {
			p.join(fmt.Sprintf("instances/%s", instance))
			if db = p.Ctx.String("database"); db != "" {
				p.join(fmt.Sprintf("databases/%s", db))
				if scheme = p.Ctx.String("table"); scheme != "" {
					p.join(fmt.Sprintf("schema/%s", scheme))
				}
			}
		}
	case "gcr":
	case "run", "functions":
		var region, name string
		if region = p.Ctx.String("region"); name != "" {
			p.join(fmt.Sprintf("details/%s", region))

			if name = p.Ctx.String("name"); name != "" {
				p.join(name)
			}
		}

		switch product {
		case "run":
		case "functions":
		}
	default:
		p.join("home/dashboard")
	}

	return
}

func (p *Provider) addProjectParam() {
	params := url.Values{}
	params.Add("project", p.SDKConfig.Core.Project)
	p.URL.RawQuery = params.Encode()
	return
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}
