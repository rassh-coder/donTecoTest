package config

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

type (
	DB struct {
		DBHost     string `env:"db_host"`
		DBPort     string `env:"db_port"`
		DBUser     string `env:"db_user"`
		DBPassword string `env:"db_pass"`
		DBName     string `env:"db_name"`
	}

	Host struct {
		Host string `env:"APP_HOST"`
		Port string `env:"APP_PORT"`
	}

	Config struct {
		DB
		Host
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg.DB = DB{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
	}
	cfg.Host = Host{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}

	return &cfg, nil
}
