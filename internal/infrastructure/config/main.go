package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	config Configuration
	once   sync.Once
)

type Configuration struct {
	Application `mapstructure:",squash"`
	Database    `mapstructure:",squash"`
}

type Application struct {
	Environment string `mapstructure:"APP_ENVIRONMENT"`
	Port        int    `mapstructure:"APP_PORT"`
}

type Database struct {
	Driver       string `mapstructure:"DB_DRIVER"`
	Host         string `mapstructure:"DB_HOST"`
	Port         string `mapstructure:"DB_PORT"`
	User         string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
	DatabaseName string `mapstructure:"DB_NAME"`
	SSLmode      string `mapstructure:"DB_SSL_MODE"`
}

func LoadConfig() Configuration {
	once.Do(func() {
		viper.SetEnvPrefix("urlify")

		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		err = viper.Unmarshal(&config)

		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		fmt.Println("DRIVER: " + config.Driver)
	})

	return config
}
