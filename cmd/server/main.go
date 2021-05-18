package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Willsem/compare-trajectories/app/server"
	"github.com/Willsem/compare-trajectories/app/server/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.New()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
