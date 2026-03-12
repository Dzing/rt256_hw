package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Http struct {
			Addr string `yaml:"addr"`
		} `yaml:"http"`
		// Loms struct {
		// 	Addr string `yaml:"addr"`
		// } `yaml:"loms"`
		// Prod struct {
		// 	Addr  string `yaml:"addr"`
		// 	Token string `yaml:"token"`
		// } `yaml:"prod"`
	}
)

// NewConfig returns app config.
func NewConfig(cfgPath string) (*Config, error) {

	if cfgPath == "" {
		return nil, fmt.Errorf("config error: %s", "no path found")
	}

	yamlFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("config error: fail to read file %s", cfgPath)
	}

	config := &Config{}
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		return nil, fmt.Errorf("config error: fail to unmarshal from file %s", cfgPath)
	}

	return config, nil
}
