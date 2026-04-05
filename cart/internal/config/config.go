package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Srv struct {
			HttpAddr string `yaml:"http"`
			GrpcAddr string `yaml:"grpc"`
		} `yaml:"srv"`
		Loms struct {
			GrpcAddr string `yaml:"grpc"`
		} `yaml:"loms"`
		Prod struct {
			GrpcAddr string `yaml:"grpc"`
			Token    string `yaml:"token"`
		} `yaml:"prod"`
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
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, fmt.Errorf("config error: fail to unmarshal from file %s", cfgPath)
	}

	return config, nil
}
