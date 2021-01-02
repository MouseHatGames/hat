package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

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
	cfg.Dashboard.Columns = 3

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	for p, v := range cfg.Widgets {
		v.Path = p
	}

	return &cfg, nil
}

func (c *Config) OrderedWidgets() []*widget.Widget {
	all := make([]*widget.Widget, 0, len(c.Widgets))

	for _, w := range c.Widgets {
		all = append(all, w)
	}

	sort.Slice(all, func(i, j int) bool {
		return strings.Compare(all[i].Path, all[j].Path) > 0
	})

	return all
}
