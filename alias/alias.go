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

package alias

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	defaultConfigPath = os.Getenv("HOME") + "/.biko/config.toml"
)

// TomlConfig ...
type TomlConfig struct {
	AWS             map[string]interface{} `toml:"aws"`
	Azure           map[string]interface{} `toml:"azure"`
	GCP             map[string]interface{} `toml:"gcp"`
	Datadog         map[string]interface{} `toml:"datadog"`
	Google          map[string]interface{} `toml:"google"`
	GoogleWorkspace map[string]interface{} `toml:"googleworkspace"`
	Youtube         map[string]interface{} `toml:"youtube"`
	PagerDuty       map[string]interface{} `toml:"pagerduty"`
	Github          map[string]interface{} `toml:"github"`
	CircleCI        map[string]interface{} `toml:"circleci"`
	Firebase        map[string]interface{} `toml:"firebase"`
	JIRA            map[string]interface{} `toml:"jira"`
}

// GetConfig ...
func GetConfig() (*TomlConfig, error) {
	confPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	conf, err := getConfig(confPath)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func getConfigPath() (string, error) {
	var confPath string
	if confPath = os.Getenv("BIKO_CONFIG_PATH"); confPath == "" {
		confPath = defaultConfigPath
	}

	if !configFileExists(confPath) {
		return "", fmt.Errorf("config file doesn't exists in %s", confPath)
	}

	return confPath, nil
}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getConfig(loadPath string) (*TomlConfig, error) {
	var conf TomlConfig
	if _, err := toml.DecodeFile(loadPath, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
