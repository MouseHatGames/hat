//go:generate bindata ./dist/...

package main

import (
	"log"

	"github.com/MouseHatGames/hat-ui/config"
	"github.com/kataras/iris/v12"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	hat, err := 

	app := iris.New()
	app.Get("/api/widgets", func(ctx iris.Context) {
		ctx.JSON(cfg.WidgetRows)
	})
	app.Get("/api/data", func(ctx iris.Context) {
		data := make(map[string]interface{})

		for _, row := range cfg.WidgetRows {
			for path, w := range row {
				
			}
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
