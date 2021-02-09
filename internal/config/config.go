package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int        `yaml:"port"`
	DataPath string `yaml:"dataPath"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := Config{
		Port: 4659,
		DataPath: "./data",
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &cfg, nil
		}

		return nil, fmt.Errorf("read file: %w", err)
	}

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
