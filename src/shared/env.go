package shared

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

func GetEnv() Env {
	var e Env
	err := envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
