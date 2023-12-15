package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host         string `yaml:"Host"`
	MySqlUrl     string `yaml:"MySqlUrl"`
	MySqlMaxIdle int    `yaml:"MySqlMaxIdle"`
	MySqlMaxOpen int    `yaml:"MySqlMaxOpen"`
}

var Instance Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&Instance)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
