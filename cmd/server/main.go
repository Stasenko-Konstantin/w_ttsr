package main

import (
	"github.com/Stasenko-Konstantin/w_ttsr/internal/config"
	"github.com/Stasenko-Konstantin/w_ttsr/internal/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"os"
	"time"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "w_ttsr_repository_query_total",
		Help: "The total number of queries processed.",
	})
)

func recordMetrics() {
	go func() {
		for {
			counter.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

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
