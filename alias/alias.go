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
	GCP       map[string]interface{} `toml:"gcp"`
	Datadog   map[string]interface{} `toml:"datadog"`
	Google    map[string]interface{} `toml:"google"`
	Youtube   map[string]interface{} `toml:"youtube"`
	PagerDuty map[string]interface{} `toml:"pagerduty"`
	Github    map[string]interface{} `toml:"github"`
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
