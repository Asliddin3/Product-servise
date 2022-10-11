package config

import (
	"os"
	"github.com/spf13/cast"
)

type Config struct{
	Environment string
	PostgresHost string
	PostgresPort int
	PostgresDatabase string
	PostgresUser string
	PostgresPassword string
	LogLevel string
	RPCPort string
	ReviewServiceHost string
	ReviewServicePort int
}
func Load() Config{
	c:=Config{}
	c.Environment=cast.ToString(getOrReturnDefault("ENVIRONMENT","develop"))
	c.PostgresHost=cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	c.PostgresPort=cast.ToInt(getOrReturnDefault("POSTGRES_PORT",5432))
	c.PostgresDatabase=cast.ToString(getOrReturnDefault("POSTGRES_DATABASE","productdb"))
	c.PostgresUser=cast.ToString(getOrReturnDefault("POSTGRES_USER","postgres"))
	c.PostgresPassword=cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","compos1995"))

	c.ReviewServiceHost=cast.ToString(getOrReturnDefault("STORE_HOST","localhost"))
	c.ReviewServicePort=cast.ToInt(getOrReturnDefault("STORE_PORT",8000))
	c.LogLevel=cast.ToString(getOrReturnDefault("LOG_LEVEL","debug"))

	c.RPCPort=cast.ToString(getOrReturnDefault("RPC_PORT",":9000"))
	return c
}

func getOrReturnDefault(key string,defaulValue interface{}) interface{}{
	_,exists:=os.LookupEnv(key)
	if exists{
		return os.Getenv(key)
	}
	return defaulValue
}
