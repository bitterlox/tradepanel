package main

import (
	"github.com/bitterlox/tradepanel/server/remote"
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[MAIN] ", log.LstdFlags)

type Config struct {
	Rpc remote.Config
}

func parseConfig(path string) (*Config, error) {

	var cfg Config

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {

	cfg, err := parseConfig("config.toml")
	if err != nil {
		log.Fatal("error reading config: ", err)
	}

	logger.Printf("Printing config: %+v", cfg)

	os.Exit(1)
}
