package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MouseHatGames/hat/pkg/client"
)

var ErrInvalidJson = errors.New(`invalid JSON value. If you want you store a string make sure you quote it accordingly, for example '"foo"' on bash or """foo""" on Windows`)

type withPath struct {
	Path string `arg required help:"Dot-separated path to the value to fetch"`
}

func (p *withPath) PathParts() []string {
	return strings.Split(p.Path, ".")
}

type GetCmd struct {
	withPath
}

func (c *GetCmd) Run(cl client.Client) error {
	val := cl.Get(c.PathParts()...)
	if val.Error() != nil {
		return val.Error()
	}

	fmt.Println(val.Raw())
	return nil
}

type SetCmd struct {
	withPath
	Value string `arg required help:"JSON-encoded value"`
}

func (c *SetCmd) Run(cl client.Client) error {
	err := cl.SetRaw(c.Value, c.PathParts()...)
	if err != nil {
		return err
	}

	if !json.Valid([]byte(c.Value)) {
		return ErrInvalidJson
	}

	return nil
}

type DelCmd struct {
	withPath
}

func (c *DelCmd) Run(cl client.Client) error {
	err := cl.Del(c.PathParts()...)
	if err != nil {
		return err
	}

	return nil
}

type ImportCmd struct {
	JSONPath string `arg required help:"path to the json file to import"`
	Root     string `short:"r" help:"path to import the file into"`
}

func (c *ImportCmd) Run(cl client.Client) error {
	var data interface{}

	root := strings.Split(c.Root, ".")
	if len(root) == 1 && root[0] == c.Root {
		root = nil
	}

	f, err := os.ReadFile(c.JSONPath)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err := json.Unmarshal(f, &data); err != nil {
		return fmt.Errorf("unmarshal json: %w", err)
	}

	m, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("json file must be a dictionary")
	}

	return importMap(m, root, cl)
}

func importMap(m map[string]interface{}, prefix []string, cl client.Client) error {
	for k, v := range m {
		path := append(prefix, k)

		if m, ok := v.(map[string]interface{}); ok {
			err := importMap(m, path, cl)
			if err != nil {
				return fmt.Errorf("set map at %v: %w", path, err)
			}
		}

		fmt.Printf("set %s\n", strings.Join(path, "."))

		err := cl.Set(v, path...)
		if err != nil {
			return fmt.Errorf("set value at %v: %w", path, err)
		}
	}

	return nil
}
