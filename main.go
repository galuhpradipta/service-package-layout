package main

import (
	"log"

	"github.com/galuhpradipta/service-package-layout/config"
	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

func init() {
	viper.SetConfigFile(`config.toml`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	config.EchoRouter().Run()
}
