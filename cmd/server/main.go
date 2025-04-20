package main

import (
	"github.com/Stasenko-Konstantin/w_ttsr/internal/config"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/server"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.New(wd)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(cfg.Server.Port, cfg.Pg.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s.Start())
}
