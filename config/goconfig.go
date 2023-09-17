package config

import (
	"errors"
	"os"

	godotenv "github.com/joho/godotenv"
	errs "github.com/prcryx/rss_aggregator/common/errs"
)

func LoadConfig() (EnvVars, error) {
	var env EnvVars
	godotenv.Load(".env")

	env = EnvVars{
		Port: os.Getenv("PORT"),
	}
	if env.Port == "" {
		return env, errors.New(errs.PortNotFound)
	}

	return env, nil

}
