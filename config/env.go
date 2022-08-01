package config

type config struct {
	Port string `env:"PORT" envDefault:"8080"`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`

	NaverClientID     string `env:"NAVER_CLIENT_ID" envDefault:""`
	NaverClientSecret string `env:"NAVER_CLIENT_SECRET" envDefault:""`
}
