package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpServer struct {
		Host string `envconfig:"SERVER_HOST"`
		Port string `envconfig:"SERVER_PORT" default:":8070"`
	}

	GrpcServer struct {
		Host string `envconfig:"SERVER_HOST"`
		Port string `envconfig:"SERVER_PORT" default:":9000"`
	}


	PostgresDB struct {
		DbString string `required:"true" envconfig:"POSTGRES_DB_STRING" default:"user=postgres password=password dbname=postgresUserdb port=5430 sslmode=disable"` // "postgres://postgres:password@localhost:5430/postgresUserdb?sslmode=disable"`
		URL      string `required:"true" envconfig:"POSTGRES_DB_URL" default:"localhost:5430"`
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