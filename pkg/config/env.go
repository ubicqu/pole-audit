package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type properties struct {
	DSN string `env:"DSN" envDefault:"root:@tcp(localhost:3306)"`
	ENV string `env:"ENV" envDefault:"dev"`
}

func (p properties) IsDev() bool {
	return Properties.ENV == "dev"
}

var Properties properties

func init() {
	if err := env.Parse(&Properties); err != nil {
		panic(err)
	}
	if Properties.IsDev() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
