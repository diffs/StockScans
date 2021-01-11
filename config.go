package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type ConfigDatabase struct {
	TickersFile     string `yaml:"tickersFile"`
	AlpacaKeyID     string `yaml:"alpacaKeyId"`
	AlpacaSecretKey string `yaml:"alpacaSecretKey"`
	UserAgent       string `yaml:"userAgent"`
	Deviance        float64 `yaml:"deviance"`
}

var cfg ConfigDatabase

func loadConfig() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "cannot find") {
			generateDefaultConfig()
		}
	}
}

func generateDefaultConfig() {
	genned := ConfigDatabase{
		TickersFile: "tickers.txt",
		AlpacaKeyID: "ALPACA_KEY_ID_HERE",
		AlpacaSecretKey: "ALPACA_SECRET_KEY_HERE",
		UserAgent: "Custom Stock Scanner Utility",
		Deviance: 0.0069,
	}
	mar, err := yaml.Marshal(genned)
	if err != nil {
		fmt.Println("Error marshalling yml")
	}
	err = ioutil.WriteFile("config.yml", mar, 0644)
	if err != nil {
		fmt.Println("Error writing default config!")
	}

	fmt.Println("Default config file written! Go edit it")
	os.Exit(0)

}

