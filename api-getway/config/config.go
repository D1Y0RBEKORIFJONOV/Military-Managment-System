package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	SoldierServiceHost string
	SoldierServicePort int

	StorehouseServiceHost string
	StorehouseServicePort int

	GroupServiceHost string
	GroupServicePort int

	CtxTimeout int

	LogLevel string
	HTTPPort string

	SignInKey         string
	AccessTokenTimout int
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":9999"))

	c.SignInKey = cast.ToString(getOrReturnDefault("SIGN_IN_KEY", "SAd2dsaSAXXcaSadSWdsHaaFs"))
	c.AccessTokenTimout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMEOUT", 6000))

	c.StorehouseServiceHost = cast.ToString(getOrReturnDefault("STOREHOUSE_SERVICE_HOST", "localhost"))
	c.StorehouseServicePort = cast.ToInt(getOrReturnDefault("STOREHOUSE_SERVICE_PORT", 9001))

	c.GroupServiceHost = cast.ToString(getOrReturnDefault("GROUP_SERVICE_HOST", "localhost"))
	c.GroupServicePort = cast.ToInt(getOrReturnDefault("GROUP_SERVICE_PORT", 9002))

	c.SoldierServiceHost = cast.ToString(getOrReturnDefault("SOLDIER_SERVICE_HOST", "localhost"))
	c.SoldierServicePort = cast.ToInt(getOrReturnDefault("SOLDIER_SERVICE_PORT", 9000))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
