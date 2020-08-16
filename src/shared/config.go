package shared

import (
	"fmt"
)

type CorConfig struct {
	AllowOrigin string
	AllowMethod string
	AllowHeaders string
	AllowExposeHeaders string
}

type WebConfig struct {
	Cors CorConfig
}

type Config struct {
	Web WebConfig
}

func GetConfig(env Env) Config {
	return Config{
		WebConfig{
			CorConfig{
				AllowOrigin:        fmt.Sprintf("%s", env.WebUrl),
				AllowHeaders:       "GET, PUT, POST, DELETE, HEAD, OPTIONS",
				AllowMethod:        "Content-Type",
				AllowExposeHeaders: "",
			},
		},
	}
}