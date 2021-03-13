package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
