//go:generate bindata ./dist/...

package main

import (
	"log"

	"github.com/MouseHatGames/hat-ui/config"
	"github.com/MouseHatGames/hat-ui/widget"
	"github.com/MouseHatGames/hat/pkg/client"
	"github.com/kataras/iris/v12"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	hat, err := client.Dial(cfg.Endpoint)
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err)
	}

	app := iris.New()
	app.Get("/api/data", func(ctx iris.Context) {
		resp := &struct {
			Widgets []map[string]*widget.Widget `json:"widgets"`
			Data    map[string]interface{}      `json:"data"`
		}{
			Widgets: cfg.WidgetRows,
			Data:    make(map[string]interface{}),
		}

		for _, row := range cfg.WidgetRows {
			for path, w := range row {
				clvalue := hat.Get(client.SplitPath(path)...)
				if err := clvalue.Error(); err != nil {
					//TODO Handle error
					continue
				}

				value, err := w.UnmarshalValue(clvalue.Raw())
				if err != nil {
					//TODO Handle error
					continue
				}

				resp.Data[path] = value
			}
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

		hat.Set(string(body), client.SplitPath(path)...)
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
