package config

import (
	"github.com/joho/godotenv"
	"os"
	"test-ihsan/model"
)

type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load("test.env")
	if err != nil {
		return err
	}
	return err
}

func (config *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) CatchError(err error) {
	if err != nil {
		panic(any(err))
	}
}

func (c *Config) GetDBConfig() model.DBConfig {
	return model.DBConfig{
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
