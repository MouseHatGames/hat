package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MouseHatGames/hat-ui/widget"
	"gopkg.in/yaml.v3"
)

type Config struct {
	WidgetRows []map[string]*widget.Widget `yaml:"widgets"`
}

func Load(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}

		return nil, fmt.Errorf("read file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	for _, row := range cfg.WidgetRows {
		for _, w := range row {
			if valid, err := w.IsValid(); !valid {
				return nil, fmt.Errorf("widget '%s': %w", w.Title, err)
			}
		}
	}

	return &cfg, nil
}
