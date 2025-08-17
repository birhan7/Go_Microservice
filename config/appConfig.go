package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HostAddress string
}

func EnvSetup() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return AppConfig{}, errors.New("unable to load environment variables")
	}
	host_address := os.Getenv("HOST_ADDRESS")
	return AppConfig{HostAddress: host_address}, nil
}
