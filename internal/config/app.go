package config

type AppConfig struct {
	Service string `env:"SERVICE,required"`
	Env     string `env:"ENV,required"`
}
