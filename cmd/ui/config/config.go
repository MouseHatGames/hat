package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/MouseHatGames/hat-ui/widget"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Endpoint    string                    `yaml:"endpoint"`
	WidgetsNode yaml.Node                 `yaml:"widgets"`
	Widgets     map[string]*widget.Widget `yaml:"-"`

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

	cfg.Widgets, err = parseWidgets(cfg.WidgetsNode)
	if err != nil {
		return nil, fmt.Errorf("parse widgets: %w", err)
	}

	return &cfg, nil
}

func parseWidgets(node yaml.Node) (map[string]*widget.Widget, error) {
	widgets := make(map[string]*widget.Widget)

	i := 0
	var path string
	for _, node := range node.Content {
		if node.Kind == yaml.ScalarNode {
			path = node.Value
		} else {
			w := new(widget.Widget)
			if err := node.Decode(w); err != nil {
				return nil, fmt.Errorf("parse widget '%s': %w", path, err)
			}

			w.Index = i
			w.Path = path
			widgets[path] = w
			i++
		}
	}

	return widgets, nil
}

func (c *Config) OrderedWidgets() []*widget.Widget {
	all := make([]*widget.Widget, 0, len(c.Widgets))

	for _, w := range c.Widgets {
		all = append(all, w)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].Index < all[j].Index
	})

	return all
}
