package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   struct {
		Host string `envconfig:"SERVER_HOST"`
		Port string `envconfig:"SERVER_PORT" default:":8070"`
	}


	PostgresDB struct {
		DbString string `required:"true" envconfig:"POSTGRES_DB_STRING" default:"user=postgres password=password dbname=postgres sslmode=disable"`
		URL      string `required:"true" envconfig:"POSTGRES_DB_URL" default:"localhost:8080"`
	}
}

func LoadConfig() *Config {
	var cnf Config
	err := envconfig.Process("", &cnf)
	if err != nil {
		fmt.Printf("InitConfig: Error %s\n", err.Error())
	}

	return &cnf
}