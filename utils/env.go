package utils

import "github.com/caarlos0/env/v6"

type ConfigurationEnv struct {
	SigningKey string `env:"SIGNING_KEY" envDefault:"mybasesigingkey"`
	DSN        string `env:"DSN" envDefault:"root:root@tcp(127.0.0.1:3306)/helloibe"`
}

var Config ConfigurationEnv

func LoadEnv() {
	if err := env.Parse(&Config); err != nil {
		panic(err.Error())
	}

}
