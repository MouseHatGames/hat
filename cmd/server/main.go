package main

import (
	"flag"
	"log"

	"github.com/MouseHatGames/hat/internal/config"
	"github.com/MouseHatGames/hat/internal/server"
	"github.com/MouseHatGames/hat/internal/store"
)

func main() {
	configPath := flag.String("config", "config.yaml", "")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	store, err := store.NewStore("store.db")
	if err != nil {
		log.Fatalf("failed to create store: %s", err)
	}
	defer store.Close()

	if err := server.Start(config, store); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
