package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vkunssec/golang-mongodb-logger/pkg/core"
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
