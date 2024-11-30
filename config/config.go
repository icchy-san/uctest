package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBName     string `envconfig:"DB_NAME" default:"invoice"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"password"`
}

func ReadFromEnv() (*Env, error) {
	var env Env

	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to read env")
	}

	return &env, nil
}
