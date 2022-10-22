package server

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT string `env:"POST"`
	HOST string `env:"HOST"`
}

func NewConfig() string {
	p, h := getEnvFromFile()
	config := Config{
		PORT: p,
		HOST: h,
	}
	return prepareAddr(&config)
}

func getEnvFromFile() (string, string) {
	viper.SetConfigFile("./../../configs/.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	port, ok := viper.Get("PORT").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	host, ok := viper.Get("HOST").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return port, host
}

func prepareAddr(config *Config) string {
	return config.HOST + ":" + config.PORT
}
