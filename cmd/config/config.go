package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the configuration struct the yml file is unmarshalled into
type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`

	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}

func LoadConfigs(configPath string) (Config, error) {
	var config Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("error: %v", err)
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("error: %v", err)
		return config, err
	}

	return config, nil
}
