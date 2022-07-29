package config

type config struct {
	Port string `env:"PORT" envDefault:"3000"`

	AuthKey    string `env:"AUTHKEY" envDefault:""`
	AuthSecret string `env:"AUTHSECRET" envDefault:""`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
}
