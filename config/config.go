package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Logging LoggingConfig `yaml:"logging"`
	Api     ApiConfig     `yaml:"api"`
}

type LoggingConfig struct {
	Level string `yaml:"level" envconfig:"LOGGING_LEVEL" default:"info"`
}

type ApiConfig struct {
	Address string `yaml:"address" envconfig:"API_ADDRESS" default:"localhost"`
	Port    uint16 `yaml:"port" envconfig:"API_PORT" default:"8080"`
}

// Singleton config instance
var globalConfig *Config

func Load(configFile string) (*Config, error) {
	cfg := &Config{}
	// Load config file as YAML if provided
	if configFile != "" {
		buf, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("error reading config file: %s", err)
		}
		err = yaml.Unmarshal(buf, cfg)
		if err != nil {
			return nil, fmt.Errorf("error parsing config file: %s", err)
		}
	}
	// Load config values from environment variables
	// This also has the side effect of setting default values from annotations
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing environment: %s", err)
	}
	// Assign to global config var
	globalConfig = cfg
	return cfg, nil
}

func GetConfig() *Config {
	return globalConfig
}
