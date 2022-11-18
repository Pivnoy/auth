package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port      string `env:"PORT" envDefault:"9000"`
	JwtSecret string `env:"SECRET" envDefault:"LEXA_KUTSENKA"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"error"`
	MongoURL  string `env:"MONGO_RUL" envDefault:"mongodb://mongo:mongo@mongodb:27017"`
	MongoDB   string `env:"MONGO_DB" envDefault:"users"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	env.Parse(cfg)
	fmt.Println(cfg)
	return cfg, nil
}
