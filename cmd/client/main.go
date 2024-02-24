package main

import (
	"fmt"
	"os"

	"github.com/MouseHatGames/hat/pkg/client"
	"github.com/alecthomas/kong"
)

type Globals struct {
	Endpoint string `short:"e" default:"127.0.0.1:4659" help:"Endpoint of the hat server. For example: 127.0.0.1:4659" placeholder:"ADDR"`
}

type CLI struct {
	Globals

	Get    GetCmd    `cmd help:"Fetches a value by its path and prints it JSON-encoded"`
	Set    SetCmd    `cmd help:"Sets a key's JSON-encoded value"`
	Del    DelCmd    `cmd help:"Deletes a key. Does not fail if the key doesn't exist"`
	Import ImportCmd `cmd help:"Imports a JSON file"`
}

func (v *Globals) AfterApply(ctx *kong.Context) error {
	cl, err := client.Dial(v.Endpoint)
	if err != nil {
		return err
	}
	ctx.BindTo(cl, (*client.Client)(nil))

	return nil
}

func main() {
	cli := CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("hat"),
		kong.UsageOnError(),
		kong.Vars{
			"version": "0.0.1",
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)

	fmt.Fprintln(os.Stderr, "OK")
}
