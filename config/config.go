package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Path string
	}
}

func LoadConfig() Config {
	var config Config

	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Невозможно декодировать в структуру: %v", err)
	}

	return config
}
