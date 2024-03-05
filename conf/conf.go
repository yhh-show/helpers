package conf

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/yhh-show/helpers/file"
	"github.com/yhh-show/helpers/jsons"
	"github.com/yhh-show/helpers/logger"
	"os"
)

func Load[T any]() (*T, error) {
	envFile, err := file.Find(".env")
	if err != nil {
		envFile, err = file.Find(".env." + os.Getenv("APP_ENV"))
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

	logger.Println("load conf:", jsons.ToString(conf))

	return conf, nil
}

func ForceLoad[T any]() *T {
	conf, err := Load[T]()
	if err != nil {
		logger.Fatalf("ForceLoad error: %v", err)
	}
	return conf
}
