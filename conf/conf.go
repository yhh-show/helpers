package conf

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/yhh-show/helpers/file"
	"github.com/yhh-show/helpers/jsons"
	"github.com/yhh-show/helpers/logger"
)

func LoadByName[T any](name string) (*T, error) {
	var envFile string
	var err error

	if name != "" {
		envName := ".env." + name
		logger.L.Println("envName:", envName)
		envFile, err = file.Find(envName)
	} else {
		err = fmt.Errorf("no env name")
	}
	if err != nil {
		envFile, err = file.Find(".env")
		if err != nil {
			return nil, fmt.Errorf("error find .env file: %w", err)
		}
	}

	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	conf := new(T)
	if err := env.Parse(conf); err != nil {
		return nil, fmt.Errorf("error parse env: %w", err)
	}

	logger.L.Println("load conf:", envFile, jsons.ToString(conf))

	return conf, nil
}

func Load[T any]() (*T, error) {
	return LoadByName[T]("")
}

func ForceLoadByName[T any](name string) *T {
	conf, err := LoadByName[T](name)
	if err != nil {
		logger.L.Fatalf("ForceLoad error: %v", err)
	}
	return conf
}

func ForceLoad[T any]() *T {
	return ForceLoadByName[T]("")
}
