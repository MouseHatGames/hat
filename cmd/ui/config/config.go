package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MouseHatGames/hat-ui/widget"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Endpoint string                    `yaml:"endpoint"`
	Widgets  map[string]*widget.Widget `yaml:"widgets"`

	Dashboard struct {
		Columns int `yaml:"columns"`
	} `yaml:"dashboard"`
}

func Load(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}

		return nil, fmt.Errorf("read file: %w", err)
	}

	cfg := Config{
		Endpoint: "127.0.0.1:4659",
	}

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	return &cfg, nil
}
