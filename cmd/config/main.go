package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	config Configuration
	once sync.Once
)

type Configuration struct {
	Application
	Database
}

type Application struct {
	Environment string `mapstructure:"APP_ENVIRONMENT"`
	Port        int `mapstructure:"APP_PORT"`
}

type Database struct {
	Driver   string `mapstructure:"DB_DRIVER_NAME"`
	Schema   string `mapstructure:"DB_SCHEMA"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Debug    string `mapstructure:"DB_DEBUG"`
}

func LoadConfig() Configuration {
	once.Do(func() {
		viper.SetEnvPrefix("urlify")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		err = viper.Unmarshal(&config)
	})

	return config
}