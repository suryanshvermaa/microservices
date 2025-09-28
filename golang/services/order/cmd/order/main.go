package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/suryanshvermaa/microservices/golang/services/order"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_URL"`
	CatalogURL  string `envconfig:"CATALOG_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
			return err
		}
		return
	})
	defer r.Close()
	log.Println("listening on port 8080...")
	s := order.NewService(r)
	// s Service, accountURL, catalogURL string, port int
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, 8080))
}
