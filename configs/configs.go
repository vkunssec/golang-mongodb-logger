package configs

import (
	"os"

	"golang-mongodb-logger/pkg/core"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	logger := core.Logger()

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error().Msg(".env file not exists")
	}
	return os.Getenv(key)
}

const (
	ServerName = "Golang Mongodb Logger Example"
)
