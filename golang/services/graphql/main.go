package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envConfig:"ACCOUNT_SERVICE_URL"`
	CatalogUrl string `envConfig:"CATALOG_SERVICE_URL"`
	OrderUrl   string `envCofig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQlServer(cfg.AccountUrl, cfg.CatalogUrl, cfg.OrderUrl)
	if err != nil {
		log.Fatal(err)
	}
	// http.Handle("/graphql",handler.New(s.))
}
