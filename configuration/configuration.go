package configuration

import (
	"errors"
	"log"
	"os"
)

type PersistenceType string

type Configuration struct {
	Port        string
	Persistence PersistenceType
}

var (
	MemoryType          PersistenceType = "MEMORY"
	DatabaseType        PersistenceType = "DATABASE"
	StaticConfiguration                 = &Configuration{Persistence: MemoryType}
)

func init() {
	InitEnv()
}

func InitEnv() {
	StaticConfiguration.Port = GetEnvValue("PORT")
	persistenceType := GetEnvValue("PERSISTENCE")
	if persistenceType == string(DatabaseType) {
		StaticConfiguration.Persistence = DatabaseType
	}
	log.Printf("Persistence used: %s\n", StaticConfiguration.Persistence)
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
