package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_CONNECTION        string `mapstructure:"DB_CONNECTION"`
	DB_HOST              string `mapstructure:"DB_HOST"`
	DB_PORT              int    `mapstructure:"DB_PORT"`
	DB_DATABASE          string `mapstructure:"DB_DATABASE"`
	DB_USERNAME          string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD          string `mapstructure:"DB_PASSWORD"`
	JWT_SECRET_KEY       string `mapstructure:"JWT_SECRET_KEY"`
	JWT_ISSUER           string `mapstructure:"JWT_ISSUER"`
	JWT_EXPIRATION_HOURS int64  `mapstructure:"JWT_EXPIRATION_HOURS"`
}

func NewConfig() *Config {
	// TODO: RESOLVE CONFIG PATH
	c, err := LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	return &c
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
