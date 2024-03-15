package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(".env file not exists")
	}
	return os.Getenv(key)
}

const (
	ServerName = "Golang Mongodb Logger Example"
)
