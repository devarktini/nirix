package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  int
	DatabaseURL string
	ENV         Enrironment
}

var ConfigInstance *Config

func LoadConfig() error {
	godotenv.Load()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // default port
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid PORT value: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://localhost:5432/mydb" // default database URL
	}

	envStr := os.Getenv("ENV")
	var env Enrironment
	switch envStr {
	case "dev":
		env = Development
	case "stage":
		env = Staging
	case "prod":
		env = Production
	default:
		env = Development // default to development
	}

	ConfigInstance = &Config{
		ServerPort:  port,
		DatabaseURL: dbURL,
		ENV:         env,
	}
	return nil
}
