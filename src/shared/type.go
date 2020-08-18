package shared

type Env struct {
	ServicePort int `envconfig:"SERVICE_PORT"`
	ApiUrl string `envconfig:"API_URL"`
	WebUrl string `envconfig:"WEB_URL"`
	RedisHost string `envconfig:"REDIS_HOST"`
}

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