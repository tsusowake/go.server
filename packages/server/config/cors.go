package config

import (
	"strings"
)

type CORSConfig struct {
	AllowOrigins string `env:"ALLOW_ORIGINS,required"`
}

func (c *CORSConfig) Decode() []string {
	origins := strings.Split(c.AllowOrigins, ",")
	for i, o := range origins {
		origins[i] = strings.TrimSpace(o)
	}
	return origins
}
