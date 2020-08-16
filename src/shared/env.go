package shared

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Env struct {
	ServicePort int `envconfig:"SERVICE_PORT"`
	ApiUrl string `envconfig:"API_URL"`
	WebUrl string `envconfig:"WEB_URL"`
}

func GetEnv() Env {
	var e Env
	err := envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
