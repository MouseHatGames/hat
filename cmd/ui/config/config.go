package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"

	"github.com/MouseHatGames/hat-ui/widget"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Endpoint    string                    `yaml:"endpoint"`
	Watch       bool                      `yaml:"watch"`
	WidgetsNode yaml.Node                 `yaml:"widgets"`
	Widgets     map[string]*widget.Widget `yaml:"-"`

	Dashboard struct {
		Columns int `yaml:"columns"`
	} `yaml:"dashboard"`
}

var paramRegex = regexp.MustCompile("{(.*?)}")

func Load(path string, out chan<- *Config) error {
	cfg, err := readFrom(path)
	if err != nil {
		return err
	}

	out <- cfg

	if cfg.Watch {
		go watch(path, out)
	}

	return nil
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

func readFrom(path string) (*Config, error) {
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

func watch(path string, out chan<- *Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to start config watcher: %s", err)
	}
	defer watcher.Close()

	watcher.Add(path)

	log.Print("watching configuration file for changes")

	for {
		select {
		case ev, ok := <-watcher.Events:
			if !ok {
				return
			}

			if ev.Op&fsnotify.Write == fsnotify.Write {
				cfg, err := readFrom(path)
				if err != nil {
					log.Printf("failed to reload configuration: %s", err)
				} else {
					out <- cfg
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("watcher error: %s", err)
		}
	}
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

			// for _, m := range paramRegex.FindAllStringSubmatch(path, -1) {
			// 	name := m[1]
			// 	w.ParamNames = append(w.ParamNames, name)
			// }

			w.Index = i
			w.Path = path
			widgets[path] = w
			i++
		}
	}

	return widgets, nil
}
