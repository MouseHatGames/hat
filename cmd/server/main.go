package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MouseHatGames/hat/internal/server"
	"github.com/MouseHatGames/hat/internal/store"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("storePath", ".")
	viper.SetDefault("port", "4659")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("hat")
	viper.AutomaticEnv()

	pflag.String("storePath", ".", "folder inside which the data will be stored")
	pflag.String("port", "4659", "port on which hat will listen")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	storePath := viper.GetString("storePath")

	os.MkdirAll(storePath, 0)

	store, err := store.NewStore(filepath.Join(storePath, "store.db"))
	if err != nil {
		log.Fatalf("failed to create store: %s", err)
	}
	defer store.Close()

	if err := server.Start(viper.GetInt("port"), store); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
