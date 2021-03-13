//go:generate bindata ./dist/...

package main

import (
	"encoding/json"
	"log"

	"github.com/MouseHatGames/hat-ui/config"
	"github.com/MouseHatGames/hat-ui/widget"
	"github.com/MouseHatGames/hat/pkg/client"
	"github.com/kataras/iris/v12"
)

func main() {
	cfgChan := make(chan *config.Config, 1)

	err := config.Load("config.yaml", cfgChan)
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	cfg := <-cfgChan
	go func() {
		for c := range cfgChan {
			log.Print("reloaded config")
			cfg = c
		}
	}()

	hat, err := client.Dial(cfg.Endpoint)
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err)
	}

	app := iris.New()
	app.Get("/api/data", func(ctx iris.Context) {
		resp := &struct {
			Widgets []*widget.Widget       `json:"widgets"`
			Columns int                    `json:"columns"`
			Title   string                 `json:"title"`
			Data    map[string]interface{} `json:"data"`
		}{
			Widgets: cfg.OrderedWidgets(),
			Columns: cfg.Dashboard.Columns,
			Title:   cfg.Dashboard.Title,
			Data:    make(map[string]interface{}, len(cfg.Widgets)),
		}

		for path, w := range cfg.Widgets {
			clvalue := hat.Get(client.SplitPath(path)...)
			if err := clvalue.Error(); err != nil {
				//TODO Handle error
				continue
			}

			value, err := w.UnmarshalValue([]byte(clvalue.Raw()))
			if err != nil {
				//TODO Handle error
				continue
			}

			resp.Data[path] = value
		}

		ctx.JSON(resp)
	})
	app.Post("/api/widget/:path/value", func(ctx iris.Context) {
		path := ctx.Params().Get("path")
		body, err := ctx.GetBody()
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		widget, ok := cfg.Widgets[path]
		if !ok {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		var value interface{}
		if err := json.Unmarshal(body, &value); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		mval, err := widget.TransformValue(value)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		if err := hat.Set(mval, client.SplitPath(path)...); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			log.Printf("failed to set value: %s", err)
		}
	})

	app.HandleDir("/", "./dist", iris.DirOptions{
		Asset:      GzipAsset,
		AssetInfo:  GzipAssetInfo,
		AssetNames: GzipAssetNames,
		AssetValidator: func(ctx iris.Context, name string) bool {
			ctx.Header("Vary", "Accept-Encoding")
			ctx.Header("Content-Encoding", "gzip")
			return true
		},
	})

	app.Listen(":8080")
}
