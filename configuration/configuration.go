package configuration

import (
	"errors"
	"os"
)

type Configuration struct {
	Port string
}

var StaticConfiguration = &Configuration{}

func init() {
	InitEnv()
}

func InitEnv() {
	StaticConfiguration.Port = GetEnvValue("PORT")
}

func CheckConfiguration() error {
	if StaticConfiguration.Port == "" {
		return errors.New("PORT env must be setted")
	}
	return nil
}

func GetEnvValue(s string) string {
	v := os.Getenv(s)
	os.Setenv(s, "") // set env to empty, avoid reading in server memory this value
	return v
}
