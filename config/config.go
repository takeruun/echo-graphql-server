package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	applicationDir = "app/"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	err := godotenv.Load(os.ExpandEnv("/go/src/" + applicationDir + ".env"))
	if err != nil {
		fmt.Println("env file 読み込み出来ませんでした。")
	}
	c := new(Config)

	goMode := os.Getenv("GO_MODE")
	switch goMode {
	case "development":
		c.DB.DBName = os.Getenv("DB_NAME") + "_development"
	case "test":
		c.DB.DBName = os.Getenv("DB_NAME") + "_test"
	case "prodcution":
		c.DB.DBName = os.Getenv("DB_NAME")
	default:
		c.DB.DBName = os.Getenv("DB_NAME") + "_development"
	}

	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Username = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")

	c.Routing.Port = "3000"

	return c
}
