package config

import (
	"os"
)

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	RPCPort     string

	Context struct {
		Timeout string
	}

	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SslMode  string
	}
	SoldirsServer struct {
		Host string
		Port string
	}
}

func New() *Config {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "local")
	config.RPCPort = getEnv("RPC_PORT", ":9001")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "+_+diyor2005+_+")
	config.DB.SslMode = getEnv("POSTGRES_SSLMODE", "disable")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "solders_service")

	config.SoldirsServer.Host = getEnv("SOLDIERS_SERVER_HOST", "localhost")
	config.SoldirsServer.Port = getEnv("SOLDIERS_SERVER_PORT", "9000")

	return &config
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
