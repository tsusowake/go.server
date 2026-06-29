package config

import "fmt"

type RedisConfig struct {
	Host     string `env:"REDIS_HOST,required"`
	Port     uint16 `env:"REDIS_PORT,required"`
	Password string `env:"REDIS_PASSWORD,required"`
}

func (r *RedisConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
