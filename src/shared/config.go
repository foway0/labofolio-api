package shared

import (
	"fmt"
)

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